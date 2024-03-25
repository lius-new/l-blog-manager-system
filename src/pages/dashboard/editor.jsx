import React, { useState } from "react";
import ReactQuill from "react-quill";
import { Input, Select, Option, Button } from "@material-tailwind/react";
import "react-quill/dist/quill.snow.css";
import { useEffect } from "react";

export function Editor() {
  const [value, setValue] = useState("");

  useEffect(() => {
    console.log(value);
  }, [value]);

  return (
    <div className="flex flex-col justify-center py-4 gap-y-4 gap-x-4">
      <div className="h-12 flex gap-4 items-center justify-between flex-wrap">
        <div className="flex gap-12 h-full">
          <Input label="Post Title" />
          <Select label="Select Tag">
            <Option>Material Tailwind HTML</Option>
            <Option>Material Tailwind React</Option>
            <Option>Material Tailwind Vue</Option>
            <Option>Material Tailwind Angular</Option>
            <Option>Material Tailwind Svelte</Option>
          </Select>
        </div>
        <div className="flex gap-x-2">
          <Button>暂存</Button>
          <Button>上传</Button>
        </div>
      </div>
      <ReactQuill
        className="flex-auto h-[64rem] border"
        theme="snow"
        value={value}
        onChange={setValue}
      />
    </div>
  );
}

export default Editor;
