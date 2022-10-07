import { useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/auth";

export function HomePage() {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      navigate("/todo");
    } else {
      navigate({
        pathname: "/auth/login",
        search: `?next=${encodeURI(pathname)}`,
      });
    }
  }, [isAuthenticated, pathname, navigate]);

  return null;
}
