import { ReactNode, createContext, useMemo, useState } from "react";
import * as amplitude from "@amplitude/analytics-browser";

type User = {
  id: string | null;
  name: string | null;
  email: string | null;
  role: string | null;
};

type UserContextType = {
  user: User;
  setUserDetails: (
    id: string,
    name: string,
    email: string,
    role: string,
  ) => void;
  clearUserDetails: () => void;
};

const defaultUser: User = {
  id: null,
  name: null,
  email: null,
  role: null,
};

const defaultUserContext: UserContextType = {
  user: defaultUser,
  setUserDetails: () => {},
  clearUserDetails: () => {},
};

export const UserContext = createContext<UserContextType>(defaultUserContext);

type Props = {
  children: ReactNode;
};

export const UserProvider = ({ children }: Props) => {
  const [user, setUser] = useState<User>({
    id: localStorage.getItem("userId"),
    name: localStorage.getItem("userName"),
    email: localStorage.getItem("userEmail"),
    role: localStorage.getItem("userRole"),
  });

  const setUserDetails = (
    id: string,
    name: string,
    email: string,
    role: string,
  ) => {
    localStorage.setItem("userId", id);
    localStorage.setItem("userName", name);
    localStorage.setItem("userEmail", email);
    localStorage.setItem("userRole", role);
    setUser({ id, name, email, role });
    amplitude.setUserId(email);
  };

  const clearUserDetails = () => {
    setUser(defaultUser);
    amplitude.reset();
  };

  const contextValue = useMemo(
    () => ({
      user,
      setUserDetails,
      clearUserDetails,
    }),
    [user],
  );

  return (
    <UserContext.Provider value={contextValue}>{children}</UserContext.Provider>
  );
};
