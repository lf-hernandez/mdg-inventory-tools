export async function fetchJson({
  url,
  options,
  includeAuth = true,
}: {
  url: string;
  options?: RequestInit;
  includeAuth?: boolean;
}) {
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
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  return response.json();
}
