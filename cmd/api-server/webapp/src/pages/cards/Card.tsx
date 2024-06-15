import { LockClosedIcon, LockOpenIcon } from "@heroicons/react/16/solid";
import { useState } from "react";

type Props = {
  card: Card;
};
export default function Card({ card }: Props) {
  const [answerLocked, setAnswerLocked] = useState(true);
  return (
    <div className="border rounded-lg shadow p-4 mb-4">
      <div>{card.QuestionText}</div>

      <div className="mt-2">
        <div
          onClick={() => setAnswerLocked((prev) => !prev)}
          className="flex items-center gap-1 select-none cursor-pointer hover:bg-gray-200 p-2 rounded-lg"
        >
          <div>
            {answerLocked && <LockClosedIcon className="w-4" />}
            {!answerLocked && <LockOpenIcon className="w-4" />}
          </div>
          <div>
            {" "}
            {answerLocked && <>Show Answer</>}
            {!answerLocked && <>Hide Answer</>}
          </div>
        </div>
        {!answerLocked && <div>{card.AnswerText}</div>}
      </div>
      {!answerLocked && <div className="mt-2">ok / nem ok</div>}
    </div>
  );
}
