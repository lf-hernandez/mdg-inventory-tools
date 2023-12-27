import { ApiUser } from "../types";
import { fetchJson } from "../utils/http";

const BASE_URL = `${import.meta.env.VITE_API_URL}/api`;

export const AuthService = {
  async login(
    email: string,
    password: string,
  ): Promise<{ token: string; user: ApiUser }> {
    const data = await fetchJson({
      url: `${BASE_URL}/login`,
      options: {
        method: "POST",
        body: JSON.stringify({ email, password }),
      },
      includeAuth: false,
    });
    return data;
  },

  async signup(
    name: string,
    email: string,
    password: string,
  ): Promise<{ token: string; user: ApiUser }> {
    const data = await fetchJson({
      url: `${BASE_URL}/signup`,
      options: {
        method: "POST",
        body: JSON.stringify({ name, email, password }),
      },
      includeAuth: false,
    });

    return data;
  },

  async updatePassword(
    currentPassword: string,
    newPassword: string,
  ): Promise<{ message: string }> {
    const data = await fetchJson({
      url: `${BASE_URL}/update-password`,
      options: {
        method: "POST",
        body: JSON.stringify({ currentPassword, newPassword }),
        headers: {
          "Content-Type": "application/json",
        },
      },
      includeAuth: true,
    });
    return data;
  },
};
