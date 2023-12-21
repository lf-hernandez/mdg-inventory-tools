import { ReactNode, createContext, useMemo, useState } from "react";

type User = {
  id: string | null;
  name: string | null;
  email: string | null;
};

type UserContextType = {
  user: User;
  setUserDetails: (id: string, name: string, email: string) => void;
  clearUserDetails: () => void;
};

const defaultUser: User = {
  id: null,
  name: null,
  email: null,
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
  });

  const setUserDetails = (id: string, name: string, email: string) => {
    localStorage.setItem("userId", id);
    localStorage.setItem("userName", name);
    localStorage.setItem("userEmail", email);
    setUser({ id, name, email });
  };

  const clearUserDetails = () => {
    setUser(defaultUser);
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
