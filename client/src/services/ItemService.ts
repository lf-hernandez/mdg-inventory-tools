import type { Item } from "../types";
import { fetchJson } from "../utils/http";

const BASE_URL = "/api/items";
const LIMIT = 10;

export const ItemService = {
  async searchItems(query: string): Promise<Item[]> {
    return fetchJson(`${BASE_URL}?search=${encodeURIComponent(query)}`);
  },
  async getItems(page: number = 1): Promise<Item[]> {
    return fetchJson(`${BASE_URL}?page=${page}&limit=${LIMIT}`);
  },
  async getItemById(id: number): Promise<Item> {
    return fetchJson(`${BASE_URL}/${id}`);
  },
  async createItem(item: Partial<Item>): Promise<Item> {
    return fetchJson(BASE_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item),
    });
  },
  async updateItem(id: number, item: Partial<Item>): Promise<Item> {
    return fetchJson(`${BASE_URL}/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(item),
    });
  },
};
