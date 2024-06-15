import { useEffect, useState } from "react";
import { useCategory } from "../../App";
import Card from "../cards/Card";
import { Button } from "../../ui/button";

export default function PracticeIndex() {
  const { category } = useCategory();
  const [card, setCard] = useState<Card | null>(null);
  const [msg, setMsg] = useState("");

  const fetchCard = async () => {
    const resp = await fetch("/api/practice", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Category: category,
      }),
    });

    if (resp.ok) {
      const body = await resp.json();
      if (body.Status === "ok") {
        setCard(body.Card);
        setMsg("");
      } else {
        setCard(null);
        setMsg(body.Status);
      }
    }
  };

  useEffect(() => {
    fetchCard();
  }, [category]);

  return (
    <div>
      {msg !== "" && (
        <div>
          <div className="text-red-500 bg-red-100 rounded-lg p-4">{msg}</div>
          <div className="mt-4">
            <Button className="cursor-pointer" color={"green"} onClick={() => fetchCard()}>
              Get Next Card!
            </Button>
          </div>
        </div>
      )}
      {card !== null && (
        <div>
          <Card
            card={card}
            edit={false}
            callback={() => {
              fetchCard();
            }}
          />
        </div>
      )}
    </div>
  );
}
