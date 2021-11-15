import { useEffect } from "react";
import { useHistory } from "react-router-dom";

// TODO -- Ideally this becomes a better NotFound page, but for
// now we're just going to redirect to the App page which will
// most likely rediect you to login
export function NotFoundPage() {
  const history = useHistory();

  useEffect(() => {
    history.push("/todo");
  });

  return null;
}
