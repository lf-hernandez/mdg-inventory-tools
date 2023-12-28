import type { Item } from "../types";
import { fetchJson } from "../utils/http";

const BASE_URL = `${import.meta.env.VITE_API_URL}/api/items`;
const LIMIT = 10;

export const ItemService = {
  async searchItems(query: string): Promise<Item[]> {
    return fetchJson({
      url: `${BASE_URL}?search=${encodeURIComponent(query)}`,
    });
  },
  async getItems(page = 1) {
    try {
      const response = await fetchJson({
        url: `${BASE_URL}?page=${page}&limit=${LIMIT}`,
      });
      if (
        response &&
        typeof response.items !== "undefined" &&
        typeof response.totalCount !== "undefined"
      ) {
        return {
          items: response.items,
          totalCount: response.totalCount,
        };
      } else {
        console.error("Unexpected response structure:", response);
        return { items: [], totalCount: 0 };
      }
    } catch (error) {
      console.error("Error fetching items:", error);
      throw error;
    }
  },
  async getItemById(id: number): Promise<Item> {
    return fetchJson({ url: `${BASE_URL}/${id}` });
  },
  async createItem(item: Partial<Item>): Promise<Item> {
    return fetchJson({
      url: BASE_URL,
      options: {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(item),
      },
    });
  },
  async updateItem(id: string, item: Partial<Item>): Promise<Item> {
    return fetchJson({
      url: `${BASE_URL}/${id}`,
      options: {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(item),
      },
    });
  },
};
