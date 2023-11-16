import csv

input_file_path = "source.csv"
output_file_path = "cleaned-source.csv"

columns_to_keep = ["Item ID", "Item Description", "Qty on H", "Last"]

with open(input_file_path, mode="r", newline="", encoding="utf-8") as infile, open(
    output_file_path, mode="w", newline="", encoding="utf-8"
) as outfile:
    reader = csv.DictReader(infile)
    writer = csv.DictWriter(
        outfile, fieldnames=["item_id", "description", "quantity", "price"]
    )

    writer.writeheader()

    for row in reader:
        quantity = int(float(row["Qty on H"]))

        new_row = {
            "item_id": row["Item ID"],
            "description": row["Item Description"],
            "quantity": quantity,
            "price": row["Last"],
        }
        writer.writerow(new_row)

print("CSV file has been cleaned: ", output_file_path)
