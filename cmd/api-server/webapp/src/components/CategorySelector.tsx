import { useState } from "react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Select } from "../ui/select";
import { PlusCircleIcon } from "@heroicons/react/16/solid";

type props = {
  onChange: (category: string) => void;
};

export default function CategorySelector({ onChange }: props) {
  const [categories, setCategories] = useState(["egy", "ketto"]);
  const [current, setCurrent] = useState("");
  const [editorOpen, setEditorOpen] = useState(false);
  const [newCategory, setNewCategory] = useState("");

  const save = () => {
    setCategories((prev) => [newCategory, ...prev])
  }

  const close = () => {
    setEditorOpen((prev) => {
        setNewCategory("");
        return !prev;
      })
  }

  return (
    <div className="flex justify-end">
      <div className="mr-2 flex gap-2">
        <div
          onClick={close}
          className="w-6 flex items-center cursor-pointer"
        >
          <PlusCircleIcon />
        </div>
        {editorOpen && (
          <>
            <div>
              <Input
                type="text"
                placeholder="new category"
                value={newCategory}
                onChange={(e) => setNewCategory(e.target.value)}
              />
            </div>
            <div>
              <Button color="green" onClick={() => {
                save();
                close();
              }}>Save</Button>
            </div>
          </>
        )}
      </div>
      <div>
        <Select
          value={current}
          className="min-w-[300px]"
          onChange={(e) => {
            setCurrent(e.target.value);
            onChange(e.target.value);
          }}
        >
          <option value="">Category</option>
          {categories.map((item) => (
            <option key={item} value={item}>
              {item}
            </option>
          ))}
        </Select>
      </div>
    </div>
  );
}
