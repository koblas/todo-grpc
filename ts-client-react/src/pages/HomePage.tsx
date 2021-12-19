import { useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/auth";

export function HomePage() {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const auth = useAuth();

  useEffect(() => {
    if (auth.isAuthenticated) {
      navigate("/todo");
    } else {
      navigate({
        pathname: "/auth/login",
        search: `?next=${encodeURI(pathname)}`,
      });
    }
  }, [auth.isAuthenticated, history, pathname]);

  return null;
}
