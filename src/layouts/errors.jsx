import { Routes, Route } from "react-router-dom";
import routes from "@/routes";

export function Errors() {
  return (
    <div className="relative min-h-screen w-full">
      <Routes>
        {routes.map(
          ({ layout, pages }) =>
            layout === "errors" &&
            pages.map(({ path, element }) => {
              console.log(path);
              return <Route exact path={path} element={element} />;
            })
        )}
      </Routes>
    </div>
  );
}

Error.displayName = "/src/layout/errors.jsx";

export default Errors;
