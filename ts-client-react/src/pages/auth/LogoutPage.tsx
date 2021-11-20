import { useEffect } from "react";
import { useHistory } from "react-router-dom";
import { useAuth } from "../../hooks/auth";

export function AuthLogoutPage() {
  const history = useHistory();
  const { mutations } = useAuth();
  const logout = mutations.useLogout();

  useEffect(() => {
    logout({
      onCompleted() {
        history.push("/auth/login");
      },
    });
  });

  return null;
}
