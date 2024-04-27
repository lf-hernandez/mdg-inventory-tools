import saveAs from "file-saver";
import { useState } from "react";

import { useAnalytics } from "../hooks/useAnalytics";
import { ItemService } from "../services/ItemService";
import { LoadingSpinner } from "../shared";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";

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
      <div>
        <form onSubmit={handleSubmit} className="flex flex-col items-center">
          <div className="relative inline-block w-full md:w-6/12 hover:drop-shadow">
            <input
              type="text"
              className="form-input overflow-x-scroll border rounded-md py-2 pl-2 pr-8 w-full"
              placeholder="Enter search by part number, serial number or description"
              required
              value={query}
              onChange={(event) => setQuery(event.target.value)}
            />
            {query && (
              <button
                type="button"
                className="absolute right-0 top-2 w-8 h-6 text-gray-400"
                onClick={handleClear}
              >
                &#10005;
              </button>
            )}
          </div>
          <div>
            {!searchResults && (
              <button
                className=" bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 mt-2 w-full sm:w-auto"
                type="submit"
              >
                Search
              </button>
            )}
            {searchResults && searchResults.length > 0 && (
              <button
                onClick={handleExport}
                className="bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 mt-2 w-full sm:w-auto"
              >
                Export
              </button>
            )}
          </div>
        </form>
      </div>

      <div>
        {isSearched && isLoading ? (
          <LoadingSpinner />
        ) : isSearched &&
          (searchResults === null || searchResults.length === 0) ? (
          <p className="text-center text-gray-600 mt-8">
            No results found. Try different keywords or check for typos.
          </p>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            {searchResults?.map((item) => (
              <ItemCard key={item.id} item={item} onUpdate={handleItemUpdate} />
            ))}
          </div>
        )}
      </div>
    </section>
  );
};
