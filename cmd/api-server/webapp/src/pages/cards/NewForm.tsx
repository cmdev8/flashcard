import {  useState } from "react";
import { Field, Label } from "../../ui/fieldset";
import { Textarea } from "../../ui/textarea";
import { Button } from "../../ui/button";
import { PlusCircleIcon, XCircleIcon } from "@heroicons/react/16/solid";

type Props = {
  category: string;
  afterCreate: () => void;
};

export default function NewCardForm({ category, afterCreate }: Props) {
  const [question, setQuestion] = useState("");
  const [questionImage, setQuestionImage] = useState("");
  const [answer, setAnswer] = useState("");
  const [answerImage, setAnswerImage] = useState("");
  const [open, setOpen] = useState(false);

  const readBlob = (e: any, setter: any) =>
    new Promise(() => {
      const reader = new FileReader();

      reader.onload = function () {
        setter(reader.result);
      };

      if (e.target && e.target.files && e.target.files.length > 0) {
        reader.readAsDataURL(e.target.files[0]);
      }
    });

  const save = async () => {
    const q = await fetch("/api/card", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Category: category,
        QuestionText: question,
        QuestionImage: questionImage,
        AnswerText: answer,
        AnswerImage: answerImage
      }),
    });

    if ((await q).ok) {
      setQuestion("");
      setAnswer("");
      setQuestionImage("");
      setAnswerImage("");
    }

    afterCreate();
  };

  return (
    <div className="border rounded-lg shadow p-4 select-none">
      <div
        onClick={() => setOpen((prev) => !prev)}
        className="flex gap-1 items-center cursor-pointer"
      >
        <div className="w-4">
          {!open && <PlusCircleIcon />}
          {open && <XCircleIcon />}
        </div>
        <div>
          New Card in <span className="font-bold">{category}</span> category
        </div>
      </div>

      {open && (
        <>
          <div className="mt-2">
            <Field>
              <Label>Question</Label>
              <Textarea
                value={question}
                onChange={(e) => setQuestion(e.target.value)}
              />
            </Field>

            <div className="mt-2 flex">
              <div>
                <input
                  type="file"
                  onChange={(e) => readBlob(e, setQuestionImage)}
                />
              </div>
              <div>
                <img src={questionImage} className="max-w-[300px]" />
              </div>
            </div>
          </div>

          <div className="mt-2">
            <Field>
              <Label>Answer</Label>
              <Textarea
                value={answer}
                onChange={(e) => setAnswer(e.target.value)}
              />
            </Field>
          </div>

          <div className="mt-2 flex">
              <div>
                <input
                  type="file"
                  onChange={(e) => readBlob(e, setAnswerImage)}
                />
              </div>
              <div>
                <img src={answerImage} className="max-w-[300px]" />
              </div>
            </div>

          <div className="mt-2 flex justify-end">
            <div>
              <Button
                onClick={save}
                className="cursor-pointer w-48"
                color={"green"}
              >
                Save
              </Button>
            </div>
          </div>
        </>
      )}
    </div>
  );
}
