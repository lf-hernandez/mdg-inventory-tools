import type { Item } from "../types";
import { fetchJson, fetchText } from "../utils/http";

const BASE_URL = `${import.meta.env.VITE_API_URL}/api/items`;
const LIMIT = 10;

let cache: Record<string, { items: Array<Item>; totalCount: number }> = {};
const getItemCacheKey = (page: number) => `page_${page}`;

export const ItemService = {
  async exportInventory(): Promise<string> {
    return fetchText({
      url: `${BASE_URL}/csv`,
    });
  },
  async exportSearch(query: string): Promise<string> {
    return fetchText({
      url: `${BASE_URL}?search=${encodeURIComponent(query)}`,
    });
  },
  async searchItems(query: string): Promise<Item[]> {
    return fetchJson({
      url: `${BASE_URL}?search=${encodeURIComponent(query)}`,
    });
  },
  async getItems(page = 1): Promise<{ items: Item[]; totalCount: number }> {
    const cacheKey = getItemCacheKey(page);

    if (cache[cacheKey]) {
      return cache[cacheKey];
    }

    try {
      const response = await fetchJson({
        url: `${BASE_URL}?page=${page}&limit=${LIMIT}`,
      });
      if (
        response &&
        typeof response.items !== "undefined" &&
        typeof response.totalCount !== "undefined"
      ) {
        cache[cacheKey] = {
          items: response.items,
          totalCount: response.totalCount,
        };
        return cache[cacheKey];
      } else {
        console.error("Unexpected response structure:", response);
        return { items: [], totalCount: 0 };
      }
    } catch (error) {
      console.error("Error fetching items:", error);
      throw error;
    }
  },
  async getItem(id: number): Promise<Item> {
    return fetchJson({ url: `${BASE_URL}/${id}` });
  },
  async createItem(item: Partial<Item>): Promise<Item> {
    try {
      const newItem = await fetchJson({
        url: BASE_URL,
        options: {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(item),
        },
      });

      cache = {};

      return newItem;
    } catch (error) {
      console.error("Error creating item:", error);
      throw error;
    }
  },

  async updateItem(id: number | string, item: Partial<Item>): Promise<Item> {
    try {
      const updatedItem = await fetchJson({
        url: `${BASE_URL}/${id}`,
        options: {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(item),
        },
      });

      cache = {};

      return updatedItem;
    } catch (error) {
      console.error("Error updating item:", error);
      throw error;
    }
  },
};
