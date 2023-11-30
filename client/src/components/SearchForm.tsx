import { useState } from "react";
import { ItemService } from "../services/ItemService";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";

export const SearchForm = () => {
  const [query, setQuery] = useState("");
  const [searchResults, setSearchResults] = useState<Array<Item>>([]);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const items = await ItemService.searchItems(query);
      setSearchResults(items);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const handleClear = () => {
    setQuery("");
    setSearchResults([]);
  };

  return (
    <section>
      <h2 className="text-2xl font-bold my-4">Search for an Item</h2>
      <form onSubmit={handleSubmit} className="mb-4">
        <div className="flex">
          <input
            type="text"
            className="form-input border p-2 rounded-l-md flex-grow"
            placeholder="Enter search by part number, serial number or description"
            required
            value={query}
            onChange={(event) => setQuery(event.target.value)}
          />
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4"
            type="submit"
          >
            Search
          </button>
          <button
            type="button"
            className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded-r-md"
            onClick={handleClear}
          >
            Clear
          </button>
        </div>
      </form>
      <div>
        {searchResults.map((item) => (
          <ItemCard key={item.id} item={item} />
        ))}
      </div>
    </section>
  );
};
