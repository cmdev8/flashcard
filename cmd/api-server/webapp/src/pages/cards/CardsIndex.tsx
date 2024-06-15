import { useEffect } from "react";
import { useCategory } from "../../App";
import NewCardForm from "./NewForm";

export default function CardsIndex() {
  const { category } = useCategory();

  const fetchCards = async () => {

  }

  useEffect(() => {
    fetchCards()
  }, []);

  if(category === '') {
    return <div className="text-center m-16 text-cyan-600">
      Choose category!
    </div>
  }

  return (
    <>
      <NewCardForm category={category} afterCreate={fetchCards} />

      <div className="mt-4">todo: kartya lista</div>
    </>
  );
}
