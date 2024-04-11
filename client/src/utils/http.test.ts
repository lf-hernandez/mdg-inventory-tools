import { FetchArgs } from "../types";
import { fetchJson, fetchText } from "./http";

global.fetch = vi.fn(
  () =>
    Promise.resolve({
      ok: true,
      status: 200,
      json: () => Promise.resolve({}),
      text: () => Promise.resolve(""),
    }) as Promise<Response>,
);

describe("fetchJson", () => {
  it("should fetch JSON data with proper headers", async () => {
    const args: FetchArgs = {
      url: "https://example.com/api",
      options: { headers: { "Content-Type": "application/json" } },
    };
    await fetchJson(args);
    expect(fetch).toHaveBeenCalledWith("https://example.com/api", {
      headers: { "content-type": "application/json" },
    });
  });
});

describe("fetchText", () => {
  it("should fetch text data with proper headers", async () => {
    const args: FetchArgs = {
      url: "https://example.com/api",
      options: { headers: { "Content-Type": "text/plain" } },
    };
    await fetchText(args);
    expect(fetch).toHaveBeenCalledWith("https://example.com/api", {
      headers: { "content-type": "text/plain" },
    });
  });
});
