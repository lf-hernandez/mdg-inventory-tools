import { fetchJson } from "../utils/http";

const BASE_URL = "/api";

export const AuthService = {
  async login(email: string, password: string): Promise<string> {
    const data = await fetchJson({
      url: `${BASE_URL}/login`,
      options: {
        method: "POST",
        body: JSON.stringify({ email, password }),
      },
      includeAuth: false,
    });
    return data.token;
  },

  async signup(name: string, email: string, password: string): Promise<string> {
    const data = await fetchJson({
      url: `${BASE_URL}/signup`,
      options: {
        method: "POST",
        body: JSON.stringify({ name, email, password }),
      },
      includeAuth: false,
    });

    return data.token;
  },
};
