import { ReactNode, createContext, useMemo, useState } from "react";

type AuthContextType = {
  token: string | null;
  onLogin: (newToken: string) => void;
  onLogout: () => void;
};

const defaultAuthContext: AuthContextType = {
  token: null,
  onLogin: () => {},
  onLogout: () => {},
};

export const AuthContext = createContext<AuthContextType>(defaultAuthContext);

type Props = {
  children: ReactNode;
};

export const AuthProvider = ({ children }: Props) => {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token"),
  );

  const onLogin = (newToken: string) => {
    localStorage.setItem("token", newToken);
    setToken(newToken);
  };

  const onLogout = () => {
    localStorage.clear();
    setToken(null);
  };

  const contextValue = useMemo(() => ({ token, onLogin, onLogout }), [token]);

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};
