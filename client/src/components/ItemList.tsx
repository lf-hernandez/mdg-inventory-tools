import { saveAs } from "file-saver";
import { useEffect, useState } from "react";

import { useAnalytics } from "../hooks/useAnalytics";
import { ItemService } from "../services/ItemService";
import { LoadingSpinner } from "../shared";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";
import { PaginationControls } from "./PaginationControls";

export const ItemList = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalItems, setTotalItems] = useState(0);
  const [isLoading, setIsLoading] = useState(false);

  const { trackEvent } = useAnalytics();

  const totalPages = Math.ceil(totalItems / 10);

  const handlePageSelect = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    setCurrentPage(currentPage + 1);
  };

  const handleItemUpdate = (updatedItem: Item) => {
    setItems((currentItems) =>
      currentItems.map((item) =>
        item.id === updatedItem.id ? updatedItem : item,
      ),
    );
  };

  const handleExport = async () => {
    try {
      const response = await ItemService.exportInventory();
      const blob = new Blob([response], { type: "text/csv" });

      const currentDate = new Date().toLocaleDateString().replace("//g", "-");

      saveAs(blob, `mdg_inventory_${currentDate}.csv`);
      trackEvent("Inventory Exported");
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    const fetchItems = async () => {
      setIsLoading(true);
      try {
        const response = await ItemService.getItems(currentPage);
        if (response && response.items) {
          setItems(response.items);
          setTotalItems(response.totalCount);
        } else {
          console.error("Invalid response:", response);
        }
      } catch (error) {
        console.error(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchItems();

    const sectionTitle = document.querySelector("#itemsListSection");
    if (sectionTitle && currentPage !== 1) {
      sectionTitle.scrollIntoView({ behavior: "smooth" });
    }
  }, [currentPage]);

  return (
    <section id="itemsListSection">
      <div className="flex justify-end items-center">
        {items && items.length > 0 && (
          <button
            onClick={handleExport}
            className="bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 mb-4 rounded w-full sm:w-auto"
          >
            Export
          </button>
        )}
      </div>
      {isLoading && <LoadingSpinner />}
      {items && items.length > 0 && (
        <>
          <div id="itemsList">
            {items.map((item) => (
              <ItemCard key={item.id} item={item} onUpdate={handleItemUpdate} />
            ))}
          </div>
          <PaginationControls
            currentPage={currentPage}
            totalPages={totalPages}
            onPreviousPage={handlePreviousPage}
            onNextPage={handleNextPage}
            onPageSelect={handlePageSelect}
          />
        </>
      )}
    </section>
  );
};
