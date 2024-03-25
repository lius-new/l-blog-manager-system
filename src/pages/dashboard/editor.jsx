import React, { useState } from "react";
import ReactQuill from "react-quill";
import { Input, Select, Option, Button } from "@material-tailwind/react";
import "react-quill/dist/quill.snow.css";
import { getFileContent } from "@/libs/utils";

export function Editor() {
  const [value, setValue] = useState("");

  const [imageUploads, setImageUploads] = React.useState([]);
  const [coverImages, setCoverImages] = useState([]);
  const uploadImage = () => {
    let inputFile = document.createElement(`input`);
    inputFile.accept = "image/*";
    inputFile.type = "file";
    inputFile.click();
    inputFile.addEventListener("change", async () => {
      if (inputFile.files && inputFile.files.length > 0) {
        let fileContent = await getFileContent(inputFile.files[0]);
        if (imageUploads.findIndex((item) => item == fileContent) == -1) {
          setCoverImages([...coverImages, inputFile.files[0]]);
          setImageUploads([...imageUploads, fileContent]);
        }
      }
    });
  };

  return (
    <div className="flex flex-col justify-center py-4 gap-y-4 gap-x-4">
      <div className="h-48 lg:h-24 2xl:h-12 flex gap-4 items-center justify-between flex-wrap">
        <div className="flex gap-4 lg:gap-12 w-full 2xl:w-1/3 flex-wrap lg:flex-nowrap">
          <Input label="Post Title" />
          <Select label="Select Tag">
            <Option>Material Tailwind HTML</Option>
            <Option>Material Tailwind React</Option>
            <Option>Material Tailwind Vue</Option>
            <Option>Material Tailwind Angular</Option>
            <Option>Material Tailwind Svelte</Option>
          </Select>
        </div>
        <div className="flex gap-2 flex-wrap">
          <div className="flex gap-x-4 flex-wrap">
            {imageUploads.map((item, index) => (
              <Button key={index} className="p-0">
                <img
                  src={item}
                  className="w-10 h-10 rounded-md object-cover"
                ></img>
              </Button>
            ))}

            <Button
              className="whitespace-nowrap"
              disabled={!(coverImages.length < 2)}
              onClick={uploadImage}
            >
              Upload Cover Image
            </Button>
          </div>
          <Button>暂存</Button>
          <Button>上传</Button>
        </div>
      </div>
      <ReactQuill
        className="flex-auto h-[24rem]  lg:h-[62rem] border"
        theme="snow"
        value={value}
        onChange={setValue}
      />
    </div>
  );
}

export default Editor;
