import { useState } from "react";
import { NavLink, Outlet } from "react-router-dom";

function App() {
  return (
    <>
      <Heading />
    </>
  );
}

function Heading() {
  const menuClasses = `text-gray-700`;
  const activeFn = ({
    isActive,
    isPending,
  }: {
    isActive: boolean;
    isPending: boolean;
  }) =>
    isPending
      ? menuClasses
      : isActive
      ? `text-indigo-600 ${menuClasses}`
      : menuClasses;

  return (
    <>
      <div className="mt-4 mr-12 flex justify-between">
        <div className="flex flex-wrap items-center gap-6 px-4 sm:flex-nowrap sm:px-6 lg:px-8">
          <h1 className="text-base font-semibold leading-7 text-gray-900">
            FlashCard
          </h1>
          <div className="order-last flex w-full gap-x-8 text-sm font-semibold leading-6 sm:order-none sm:w-auto sm:border-l sm:border-gray-200 sm:pl-6 sm:leading-7">
            <NavLink to="/practice" className={activeFn}>
              <>Practice</>
            </NavLink>
            <NavLink to="/cards" className={activeFn}>
              <>Cards</>
            </NavLink>
          </div>
        </div>
        <div>
          -- category selector --
        </div>
      </div>

      <div className="m-8">
        <Outlet />
      </div>
    </>
  );
}

export default App;
