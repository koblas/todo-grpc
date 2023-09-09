import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useEffect, useMemo } from "react";
import { z } from "zod";
import * as rpcTeam from "../../rpc/team";
import { useAuth } from "../auth";
import { newFetchClient } from "../../rpc/utils";
import { useWebsocketUpdates } from "../../rpc/websocket";
import { buildCallbacksTyped, buildCallbacksTypedQuery } from "../../rpc/utils/helper";
import { RpcMutation, RpcOptions } from "../../rpc/errors";
import { must } from "../../util/assert";

const CACHE_KEY = ["team"];

export function useTeams(): rpcTeam.TeamMemberT[] {
  const { token } = useAuth();
  const client = newFetchClient({ token });
  const queryClient = useQueryClient();

  const result = useQuery<rpcTeam.TeamListResponseT>(CACHE_KEY, () => client.POST("/v1/team/team_list", {}), {
    staleTime: Infinity,
    suspense: true,
    enabled: true,
    ...buildCallbacksTypedQuery(queryClient, rpcTeam.TeamListResponse, {}),
  });

  return must(result.data?.teams);
}

export function useTeamMutations() {
  const { token } = useAuth();
  const queryClient = useQueryClient();

  const mutations = useMemo(() => {
    const client = newFetchClient({ token });

    return {
      useInviteUser(): RpcMutation<rpcTeam.TeamInviteRequestT, rpcTeam.TeamInviteResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcTeam.TeamInviteRequestT) => client.POST<rpcTeam.TeamInviteResponseT>("/v1/team/team_invite", data),
          {},
        );

        function action(data: rpcTeam.TeamInviteRequestT, handlers?: RpcOptions<rpcTeam.TeamInviteResponseT>) {
          mutation.mutate(data, buildCallbacksTyped(queryClient, rpcTeam.TeamInviteResponse, handlers));
        }

        return [action, { loading: mutation.isLoading }];
      },
      useTeamCreate(): RpcMutation<rpcTeam.TeamCreateRequestT, rpcTeam.TeamCreateResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcTeam.TeamCreateRequestT) => client.POST<rpcTeam.TeamCreateResponseT>("/v1/team/team_create", data),
          {},
        );

        function action(data: rpcTeam.TeamCreateRequestT, handlers?: RpcOptions<rpcTeam.TeamCreateResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(queryClient, rpcTeam.TeamCreateResponse, {
              ...(handlers ?? {}),
              onCompleted(result, variables) {
                queryClient.setQueriesData<rpcTeam.TeamListResponseT>(CACHE_KEY, (old) => {
                  if (!old) {
                    return undefined;
                  }

                  return { teams: [...old.teams, result.team] };
                });

                handlers?.onCompleted?.(result, variables);
              },
            }),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
      useTeamDelete(): RpcMutation<rpcTeam.TeamDeleteRequestT, rpcTeam.TeamDeleteResponseT> {
        const mutation = useMutation(
          CACHE_KEY,
          (data: rpcTeam.TeamDeleteRequestT) => client.POST<rpcTeam.TeamDeleteResponseT>("/v1/team/team_delete", data),
          {},
        );

        function action(data: rpcTeam.TeamDeleteRequestT, handlers?: RpcOptions<rpcTeam.TeamDeleteResponseT>) {
          mutation.mutate(
            data,
            buildCallbacksTyped(queryClient, rpcTeam.TeamInviteResponse, {
              ...(handlers ?? {}),
              onCompleted(result, variables) {
                queryClient.setQueriesData<rpcTeam.TeamListResponseT>(CACHE_KEY, (old) => {
                  if (!old) {
                    return undefined;
                  }
                  const teams = old.teams.filter((team) => team.id !== variables.team_id);

                  return { teams };
                });

                handlers?.onCompleted?.(result, variables);
              },
            }),
          );
        }

        return [action, { loading: mutation.isLoading }];
      },
    };
  }, [queryClient, token]);

  return mutations;
}

export function useTeamListener() {
  const queryClient = useQueryClient();
  const { addListener } = useWebsocketUpdates();

  useEffect(
    () =>
      addListener("message", (event: z.infer<typeof rpcTeam.TeamMemberEvent>) => {
        const data = rpcTeam.TeamMemberEvent.parse(event);

        queryClient.setQueriesData<rpcTeam.TeamListResponseT>(CACHE_KEY, (old) => {
          if (!old) {
            return old;
          }
          let { teams } = old;

          if (data.action === "delete") {
            teams = old.teams.filter((team) => team.id !== data.object_id);
          } else if (data.action === "update") {
            teams = old.teams.map((team) => {
              if (team.id !== event.body.id) {
                return team;
              }
              return event.body;
            });
          } else if (data.action === "create" && event.body !== null) {
            if (!old.teams.some((team) => team.id === data.body.id)) {
              teams = old.teams.concat([data.body]);
            }
          } else {
            throw new Error("Unknown Event");
          }

          return { teams };
        });
      }),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [queryClient],
  );
}
