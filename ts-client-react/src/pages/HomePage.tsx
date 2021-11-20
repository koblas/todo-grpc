import { useEffect } from "react";
import { useHistory, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/auth";

export function HomePage() {
  const history = useHistory();
  const { pathname } = useLocation();
  const auth = useAuth();

  useEffect(() => {
    if (auth.isAuthenticated) {
      history.push("/todo");
    } else {
      history.push({
        pathname: "/auth/login",
        search: `?next=${encodeURI(pathname)}`,
      });
    }
  }, [auth.isAuthenticated, history, pathname]);

  return null;
}
