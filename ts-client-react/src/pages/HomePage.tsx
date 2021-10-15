import { useEffect } from "react";
import { useHistory } from "react-router-dom";

export function HomePage() {
  const history = useHistory();

  useEffect(() => {
    history.push("/todo");
  });

  return null;
}
