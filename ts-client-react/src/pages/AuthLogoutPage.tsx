import { useEffect } from "react";
import { useHistory } from "react-router-dom";
import { useAuth } from "../hooks/auth";

export function AuthLogoutPage() {
  const history = useHistory();
  const auth = useAuth();

  useEffect(() => {
    auth.logout().then(() => {
      history.push("/auth/login");
    });
  });

  return null;
}
