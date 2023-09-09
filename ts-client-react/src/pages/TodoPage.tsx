import React from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { QueryErrorResetBoundary } from "@tanstack/react-query";
import { ErrorBoundary, FallbackProps } from "react-error-boundary";
import { Heading, Box, Text, CloseButton, Grid, Flex, Input, Button, Spinner } from "@chakra-ui/react";
import { useTodos } from "../hooks/data/todo";
import { TodoObjectT } from "../rpc/todo";
import { useAuth } from "../hooks/auth";

type FormFields = {
  text: string;
};

function ErrorView({ error, resetErrorBoundary }: FallbackProps) {
  return (
    <Box>
      <Text>{error.message}</Text>
      <Button marginLeft="5" size="md" colorScheme="blue" variant="solid" onClick={resetErrorBoundary}>
        Retry
      </Button>
    </Box>
  );
}

function Item({ todo }: { todo: TodoObjectT }) {
  const { mutations } = useTodos();
  const { task, id } = todo;

  return (
    <Box w={2 / 6} border="1px" borderColor="gray.200" margin="1" padding="2">
      <Flex justifyContent="space-between" alignItems="baseline">
        <Text>{task}</Text>
        <CloseButton size="sm" color="red.500" onClick={() => mutations.deleteTodo.mutate({ id })} />
      </Flex>
    </Box>
  );
}

function TodoList() {
  const { todos } = useTodos(true);

  return (
    <>
      {todos.map((todo) => (
        <Item key={todo.id} todo={todo} />
      ))}
    </>
  );
}

export function TodoDetail() {
  const { mutations } = useTodos();
  const { register, handleSubmit, setValue } = useForm<FormFields>();
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    navigate("/auth/login");
    return null;
  }

  function onSubmit(data: FormFields) {
    mutations.addTodo.mutate({ task: data.text });
    setValue("text", "");
  }

  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <Box w="100%" bgColor="white" p="5">
        <Heading as="h3" size="xl" textColor="gray.800" textAlign="center" fontWeight="light" padding="5">
          TODO gRPC Client
        </Heading>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Flex justifyContent="center">
            <Flex>
              <Input padding="2" width="80" placeholder="Walk my dog" type="text" {...register("text")} />
              <Button marginLeft="5" size="md" colorScheme="blue" variant="solid" onClick={handleSubmit(onSubmit)}>
                Add Todo
              </Button>
            </Flex>
          </Flex>
        </form>
      </Box>
      <Box p="5" bgColor="white">
        <Grid justifyItems="center">
          <QueryErrorResetBoundary>
            {({ reset }) => (
              <ErrorBoundary onReset={reset} FallbackComponent={ErrorView}>
                <React.Suspense fallback={<Spinner />}>
                  <TodoList />
                </React.Suspense>
              </ErrorBoundary>
            )}
          </QueryErrorResetBoundary>
        </Grid>
      </Box>
    </Box>
  );
}

export function TodoPage() {
  return (
    <Box w="100%" p="8" bgColor="gray.100">
      <TodoDetail />
    </Box>
  );
}
