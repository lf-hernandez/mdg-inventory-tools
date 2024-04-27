import { ChangeEvent, useState } from "react";
import { toast } from "react-hot-toast";

import { useAnalytics } from "../hooks/useAnalytics";
import { ItemService } from "../services/ItemService";
import type { Item } from "../types";
import { InputField } from "./InputField";

export const AddItemForm = () => {
  const [item, setItem] = useState<Partial<Item>>({
    partNumber: "",
    serialNumber: "",
    purchaseOrder: "",
    description: "",
    category: "",
    price: 0,
    quantity: 0,
  });
  const { trackEvent } = useAnalytics();

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    let updatedValue: string | number = value;

    if (name === "price") {
      updatedValue = value ? parseFloat(value) : 0;
    } else if (name === "quantity") {
      updatedValue = value ? parseInt(value, 10) : 0;
    }

    setItem((prevState) => ({
      ...prevState,
      [name]: updatedValue,
    }));
  };

  const resetForm = () => {
    setItem({
      partNumber: "",
      serialNumber: "",
      purchaseOrder: "",
      description: "",
      category: "",
      price: 0,
      quantity: 0,
      location: "",
      notes: "",
    });
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      await ItemService.createItem(item);
      resetForm();
      toast.success("Item added successfully");
      trackEvent("Inventory Added");
    } catch (e) {
      toast.error(`An error occurred trying to add item`);
      console.error("Error:", e);
      trackEvent(`Add item failed: ${e}`);
    }
  };

  return (
    <div>
      <section className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <InputField
              label="Part Number"
              name="partNumber"
              value={item.partNumber ?? ""}
              onChange={handleChange}
              required
            />
            <InputField
              label="Serial Number"
              name="serialNumber"
              value={item.serialNumber ?? ""}
              onChange={handleChange}
            />
            <InputField
              label="Purchase Order"
              name="purchaseOrder"
              value={item.purchaseOrder ?? ""}
              onChange={handleChange}
              required
            />
            <InputField
              label="Description"
              name="description"
              value={item.description ?? ""}
              onChange={handleChange}
            />
            <InputField
              label="Category"
              name="category"
              value={item.category ?? ""}
              onChange={handleChange}
            />
            <InputField
              label="Price"
              name="price"
              type="number"
              value={item.price ?? 0}
              onChange={handleChange}
            />
            <InputField
              label="Quantity"
              name="quantity"
              type="number"
              value={item.quantity ?? 0}
              onChange={handleChange}
            />
            <InputField
              label="Status"
              name="status"
              value={item.status ?? ""}
              onChange={handleChange}
            />
            <InputField
              label="Repair Order Number"
              name="repairOrderNumber"
              value={item.repairOrderNumber ?? ""}
              onChange={handleChange}
            />
            <InputField
              label="Condition"
              name="condition"
              value={item.condition ?? ""}
              onChange={handleChange}
            />
            <InputField
              label="Location"
              name="location"
              value={item.location ?? ""}
              onChange={handleChange}
            />
            <InputField
              multiline
              label="Notes"
              name="notes"
              value={item.notes ?? ""}
              onChange={handleChange}
            />
          </div>
          <div className="flex justify-end items-center mt-4">
            <button
              type="submit"
              className="rounded bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 my-4 w-full sm:w-auto"
            >
              Add Item
            </button>
          </div>
        </form>
      </section>
    </div>
  );
};
