import { saveAs } from "file-saver";
import { useEffect, useState } from "react";
import * as amplitude from "@amplitude/analytics-browser";

import { ItemService } from "../services/ItemService";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";
import { PaginationControls } from "./PaginationControls";

export const ItemList = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalItems, setTotalItems] = useState(0);
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

  useEffect(() => {
    const fetchItems = async () => {
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
      }
    };

    fetchItems();

    const sectionTitle = document.querySelector("#itemsListSection");
    if (sectionTitle) {
      sectionTitle.scrollIntoView({ behavior: "smooth" });
    }
  }, [currentPage]);

  const handleExport = async () => {
    try {
      const response = await ItemService.exportInventory();
      const blob = new Blob([response], { type: "text/csv" });

      const currentDate = new Date().toLocaleDateString().replace("//g", "-");

      saveAs(blob, `mdg_inventory_${currentDate}.csv`);
      amplitude.track("Inventory Exported");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <section id="itemsListSection">
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold my-4">Inventory</h2>
        {items && items.length > 0 && (
          <button
            onClick={handleExport}
            className="l-auto bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
          >
            Export
          </button>
        )}
      </div>
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
