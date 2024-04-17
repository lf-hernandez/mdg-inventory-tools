import { FetchArgs } from "../types";

export async function fetchJson({
  url,
  options,
  includeAuth = true,
}: FetchArgs) {
  const headers: Record<string, string> = {};

  if (options?.headers) {
    const headersInit = new Headers(options.headers);
    headersInit.forEach((value, key) => {
      headers[key] = value;
    });
  }

  if (includeAuth) {
    const token = localStorage.getItem("token");
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
  }

  const response = await fetch(url, {
    ...options,
    headers,
  });

  if (!response.ok) {
    if (response.status === 401 && window.location.pathname !== "/login") {
      localStorage.removeItem("token");
      window.location.href = "/login";
      return Promise.reject(new Error("Session expired. Please login again."));
    }
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  return response.json();
}

export async function fetchText({
  url,
  options,
  includeAuth = true,
}: FetchArgs) {
  const headers: Record<string, string> = {};

  if (options?.headers) {
    const headersInit = new Headers(options.headers);
    headersInit.forEach((value, key) => {
      headers[key] = value;
    });
  }

  if (includeAuth) {
    const token = localStorage.getItem("token");
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
  }

  const response = await fetch(url, {
    ...options,
    headers,
  });

  if (!response.ok) {
    if (response.status === 401) {
      localStorage.removeItem("token");
      window.location.href = "/login";
      return Promise.reject(new Error("Session expired. Please login again."));
    }
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.text();
}
