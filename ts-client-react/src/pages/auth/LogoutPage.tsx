import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../hooks/auth";

export function AuthLogoutPage() {
  const navigate = useNavigate();
  const { mutations } = useAuth();
  const logout = mutations.useLogout();

  useEffect(() => {
    logout({
      onCompleted() {
        navigate("/auth/login");
      },
    });
  });

  return null;
}
