import { ChangeEvent } from "react";

type Props = {
  label: string;
  name: string;
  type?: string;
  value: string | number;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
};

export const InputField = ({
  label,
  name,
  type = "text",
  value,
  onChange,
  required = false,
}: Props) => (
  <div>
    <label htmlFor={name} className="block text-sm font-medium text-gray-700">
      {label}
    </label>
    <input
      type={type}
      name={name}
      id={name}
      value={value}
      onChange={onChange}
      required={required}
      className="mt-1 focus:ring-blue-500 focus:border-blue-500 block w-full shadow-sm sm:text-sm border-gray-300"
    />
  </div>
);
