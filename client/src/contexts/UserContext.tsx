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

const defaultUserContext: UserContextType = {
  user: { id: null, name: null, email: null },
  setUserDetails: () => {},
  clearUserDetails: () => {},
};

export const UserContext = createContext<UserContextType>(defaultUserContext);

type Props = {
  children: ReactNode;
};

export const UserProvider = ({ children }: Props) => {
  const [user, setUser] = useState<User>({ id: null, name: null, email: null });

  const setUserDetails = (id: string, name: string, email: string) => {
    setUser({ id, name, email });
  };

  const clearUserDetails = () => {
    setUser({ id: null, name: null, email: null });
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
