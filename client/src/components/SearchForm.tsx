import saveAs from "file-saver";
import { useState } from "react";

import { ItemService } from "../services/ItemService";
import { LoadingSpinner } from "../shared";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";
import { useAnalytics } from "../hooks/useAnalytics";

export const SearchForm = () => {
  const [query, setQuery] = useState("");
  const [searchResults, setSearchResults] = useState<Array<Item> | null>(null);
  const [isSearched, setIsSearched] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { trackEvent } = useAnalytics();

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    setIsLoading(true);
    event.preventDefault();
    setIsSearched(true);
    if (!query.trim()) return;
    try {
      const items = await ItemService.searchItems(query);
      setSearchResults(items);
      trackEvent("Inventory Search", {
        term: query,
        result_count: items.length,
      });
    } catch (error) {
      console.error("Error:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleClear = () => {
    setQuery("");
    setSearchResults(null);
    setIsSearched(false);
    trackEvent("Search Cleared", { term: query });
  };

  const handleItemUpdate = (updatedItem: Item) => {
    setSearchResults(
      (currentResults) =>
        currentResults?.map((item) =>
          item.id === updatedItem.id ? updatedItem : item,
        ) || null,
    );
  };

  const handleExport = async () => {
    try {
      trackEvent("Export Search Results");
      const response = await ItemService.exportSearch(query);
      const blob = new Blob([response], { type: "text/csv" });

      const currentDate = new Date().toLocaleDateString().replace("//g", "-");

      saveAs(blob, `mdg_${query}_inventory_${currentDate}.csv`);
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <section>
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold my-4">Search for an item</h2>
        {searchResults && searchResults.length > 0 && (
          <button
            onClick={handleExport}
            className="l-auto bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
          >
            Export results
          </button>
        )}
      </div>
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
            type="button"
            className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4"
            onClick={handleClear}
          >
            Clear
          </button>
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-r-md"
            type="submit"
          >
            Search
          </button>
        </div>
      </form>
      <div>
        {isSearched && isLoading ? (
          <LoadingSpinner />
        ) : isSearched &&
          (searchResults === null || searchResults.length === 0) ? (
          <p className="text-center text-gray-600">
            No results found. Try different keywords or check for typos.
          </p>
        ) : (
          searchResults?.map((item) => (
            <ItemCard key={item.id} item={item} onUpdate={handleItemUpdate} />
          ))
        )}
      </div>
    </section>
  );
};
