import tabula

tabula.convert_into(
    input_path="input.pdf", output_path="output.csv", output_format="csv", pages="all"
)
