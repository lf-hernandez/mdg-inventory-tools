import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export function PublicRoute({
  component: Component,
}: {
  component: React.ComponentType;
}) {
  const { token } = useAuth();

  return !token ? <Component /> : <Navigate to="/" />;
}
