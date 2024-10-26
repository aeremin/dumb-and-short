import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter, replace,
  RouterProvider,
} from "react-router-dom";

import CreatedPage from "./CreatedPage";
import CreateLinkPage from "./CreateLinkPage";
import {resolve} from "./api";

const router = createBrowserRouter([
  {
    path: "/",
    element: <CreateLinkPage />,
  },
  {
    path: "/:id",
    loader: async ({ params }) => {
      const url = await resolve(params.id!!);
      return replace(url);
    },
  },
  {
    path: "/created/:id",
    loader: async ({ params }) => {
      return params.id!!;
    },
    element: <CreatedPage />,
  },
]);

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
