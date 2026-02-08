import React, { createContext, useContext, useEffect, useState } from "react";
import { api, setAuthToken } from "../api/client";

type AuthContextType = {
  token: string | null;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string) => Promise<void>;
  logout: () => void;
};

const AuthContext = createContext<AuthContextType>(null!);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token")
  );

  useEffect(() => {
    setAuthToken(token);
    if (token) localStorage.setItem("token", token);
    else localStorage.removeItem("token");
  }, [token]);

  async function login(email: string, password: string) {
    const res = await api.post("/auth/login", { email, password });
    setToken(res.data.token);
  }

  async function register(email: string, password: string) {
    await api.post("/auth/register", { email, password });
    await login(email, password);
  }

  function logout() {
    setToken(null);
  }

  return (
    <AuthContext.Provider value={{ token, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}
