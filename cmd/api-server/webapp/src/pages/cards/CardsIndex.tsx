import { useEffect, useState } from "react";
import { useCategory } from "../../App";
import NewCardForm from "./NewForm";
import Card from "./Card";

export default function CardsIndex() {
  const { category } = useCategory();
  const [cards, setCards] = useState<Card[]>([]);

  const fetchCards = async () => {
    const resp = await fetch("/api/card");
    if (resp.ok) {
      const body = await resp.json();
      setCards(body.cards);
    }
  };

  useEffect(() => {
    fetchCards();
  }, []);

  if (category === "") {
    return (
      <div className="text-center m-16 text-cyan-600">Choose category!</div>
    );
  }

  return (
    <>
      <NewCardForm category={category} afterCreate={fetchCards} />

      <div className="mt-4">
        {cards.map((card) => (
          <div key={card.ID}>
            <Card card={card} />
          </div>
        ))}
      </div>
    </>
  );
}
