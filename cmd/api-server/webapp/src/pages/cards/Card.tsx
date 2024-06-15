import { LockClosedIcon, LockOpenIcon } from "@heroicons/react/16/solid";
import { useState } from "react";
import { Button } from "../../ui/button";

type Props = {
  card: Card;
  callback?: () => void;
  edit: boolean;
};

export default function Card({ card, callback, edit }: Props) {
  const [answerLocked, setAnswerLocked] = useState(true);

  const sendResult = async (success: boolean) => {
    const q = await fetch("/api/result", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        CardID: card.ID,
        Success: success,
      }),
    });

    if ((await q).ok) {
      if (callback !== undefined) {
        callback();
      }
      setAnswerLocked(true);
    }
  };

  const deleteCard = async () => {
    if (!confirm("Sure?")) {
      return;
    }

    const q = await fetch(`/api/card/${card.ID}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if ((await q).ok) {
      if (callback !== undefined) {
        callback();
      }
    }
  };

  return (
    <div className="border rounded-lg shadow p-4 mb-4">
      <div>{card.QuestionText}</div>

      <div className="mt-4">
        <div
          onClick={() => setAnswerLocked((prev) => !prev)}
          className="flex items-center gap-1 select-none cursor-pointer hover:bg-gray-200 p-2 rounded-lg"
        >
          <div>
            {answerLocked && <LockClosedIcon className="w-4" />}
            {!answerLocked && <LockOpenIcon className="w-4" />}
          </div>
          <div>
            {answerLocked && <>Show Answer</>}
            {!answerLocked && <>Hide Answer</>}
          </div>
        </div>
        {!answerLocked && <div>{card.AnswerText}</div>}
      </div>
      {!answerLocked && (
        <div className="mt-4 flex gap-1">
          <div>
            <Button
              color={"red"}
              onClick={() => sendResult(false)}
              className="cursor-pointer"
            >
              Fail
            </Button>
          </div>
          <div>
            <Button
              color={"green"}
              onClick={() => sendResult(true)}
              className="cursor-pointer"
            >
              Success
            </Button>
          </div>
        </div>
      )}
      {edit && (
        <div className="mt-4 p-4 border border-red-100 rounded-lg">
          <Button
            color={"red"}
            onClick={() => deleteCard()}
            className="cursor-pointer"
          >
            Delete
          </Button>
        </div>
      )}
    </div>
  );
}
