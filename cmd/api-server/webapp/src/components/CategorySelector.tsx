import { useState } from "react";

type props = {
  onChange: (category: string) => void;
};

export default function CategorySelector({ onChange }: props) {
  const [categories, setCategories] = useState(["egy", "ketto"]);
  const [current, setCurrent] = useState("");

  return (
    <select
      value={current}
      onChange={(e) => {
        setCurrent(e.target.value);
        onChange(e.target.value);
      }}
    >
      <option value=""></option>
      {categories.map((item) => (
        <option key={item} value={item}>
          {item}
        </option>
      ))}
    </select>
  );
}
