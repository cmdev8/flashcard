import { useCategory } from "../../App";

export default function CardsIndex() {
  const { category } = useCategory();

  return (
    <>
      <div>cards index</div>
      <div>Category: {category}</div>
    </>
  );
}
