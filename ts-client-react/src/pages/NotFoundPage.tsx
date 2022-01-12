import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

// TODO -- Ideally this becomes a better NotFound page, but for
// now we're just going to redirect to the App page which will
// most likely rediect you to login
export function NotFoundPage() {
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/todo");
  });

  return null;
}
