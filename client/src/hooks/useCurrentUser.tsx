import { useContext } from "react";
import { UserContext } from "../contexts/UserContext";

export const useCurrentUser = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useCurrentUser must be used within an UserProvider");
  }
  return context;
};
