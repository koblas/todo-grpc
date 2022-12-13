import React, { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/auth";

export default function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (!isAuthenticated) {
      navigate(
        {
          pathname: "/auth/login",
          search: `?next=${encodeURI(pathname)}`,
        },
        {
          replace: true,
        },
      );
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated]);

  // eslint-disable-next-line react/jsx-no-useless-fragment
  return <>{isAuthenticated ? children : null}</>;
}
