import { useState } from "react";
import { toast } from "react-hot-toast";

import { useAnalytics } from "../hooks/useAnalytics";
import { ItemService } from "../services/ItemService";
import type { Item } from "../types";
import { InputField } from "./InputField";

type Props = {
  item: Item;
  onUpdate: (item: Item) => void;
};

export const ItemCard = ({ item, onUpdate }: Props) => {
  const [editMode, setEditMode] = useState(false);
  const [editItem, setEditItem] = useState(item);
  const { trackEvent } = useAnalytics();
  const handleEditChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    let updatedValue: string | number = value;

    if (name === "price") {
      updatedValue = value ? parseFloat(value) : 0;
    } else if (name === "quantity") {
      updatedValue = value ? parseInt(value, 10) : 0;
    }

    setEditItem((prevState) => ({
      ...prevState,
      [name]: updatedValue,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const updatedItem = await ItemService.updateItem(editItem.id, editItem);
      onUpdate(updatedItem);
      setEditMode(false);
      toast.success("Item updated successfully");
      trackEvent("Item Update", { success: true });
    } catch (error) {
      console.error("Error updating item:", error);
      toast.error("Failed to update item");
      trackEvent("Item Update", { success: false });
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="bg-white rounded-lg shadow-md p-6 mb-4 transition-all"
    >
      <div className="flex items-center justify-between mb-4">
        <h5 className="text-xl font-semibold text-gray-800">
          {`Part Number: ${item.partNumber}`}
        </h5>
        <button
          type="button"
          onClick={() => setEditMode(!editMode)}
          className="text-gray-500 hover:text-gray-700 focus:outline-none"
          aria-label={editMode ? "Close edit mode" : "Open edit mode"}
        >
          {editMode ? <span>&times; Cancel</span> : <span>&#9998; Edit</span>}
        </button>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-gray-600">
        {editMode ? (
          <>
            <InputField
              label="Description"
              type="text"
              name="description"
              value={editItem.description}
              onChange={handleEditChange}
            />
            <InputField
              label="Price"
              name="price"
              type="number"
              value={editItem.price}
              onChange={handleEditChange}
            />
            <InputField
              label="Quantity"
              name="quantity"
              type="number"
              value={editItem.quantity}
              onChange={handleEditChange}
            />
            <InputField
              label="Purchase order"
              type="text"
              name="purchaseOrder"
              value={editItem.purchaseOrder}
              onChange={handleEditChange}
            />
            <InputField
              label="Serial number"
              type="text"
              name="serialNumber"
              value={editItem.serialNumber}
              onChange={handleEditChange}
            />
            <InputField
              label="Category"
              type="text"
              name="category"
              value={editItem.category}
              onChange={handleEditChange}
            />
            <InputField
              label="Status"
              type="text"
              name="status"
              value={editItem.status}
              onChange={handleEditChange}
            />
            <InputField
              label="Repair order number"
              type="text"
              name="repairOrderNumber"
              value={editItem.repairOrderNumber}
              onChange={handleEditChange}
            />
            <InputField
              label="Condition"
              type="text"
              name="condition"
              value={editItem.condition}
              onChange={handleEditChange}
            />
            <InputField
              label="Location"
              type="text"
              name="location"
              value={editItem.location}
              onChange={handleEditChange}
            />
            <InputField
              multiline
              label="Notes"
              name="notes"
              value={editItem.notes}
              onChange={handleEditChange}
            />
          </>
        ) : (
          <>
            <p>Description: {item.description}</p>
            <p>Price: ${item.price?.toFixed(2) ?? 0.0}</p>
            <p>Quantity: {item.quantity}</p>
            <p>Purchase Order: {item.purchaseOrder}</p>
            <p>Serial Number: {item.serialNumber}</p>
            <p>Category: {item.category}</p>
            <p>Status: {item.status}</p>
            <p>Repair Order Number: {item.repairOrderNumber}</p>
            <p>Condition: {item.condition}</p>
            <p>Location: {item.location}</p>
            {item.notes && <p>Notes: {item.notes}</p>}
          </>
        )}
      </div>
      {editMode && (
        <div className="flex justify-end mt-4">
          <button
            type="submit"
            className="rounded bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 my-4 w-full sm:w-auto"
          >
            Save Changes
          </button>
        </div>
      )}
    </form>
  );
};
