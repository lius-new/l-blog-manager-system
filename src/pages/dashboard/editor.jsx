import React, { useEffect, useState } from "react";
import ReactQuill from "react-quill";
import {
  Input,
  Select,
  Option,
  Button,
  ButtonGroup,
  Chip,
  Badge,
  Dialog,
  DialogHeader,
  DialogBody,
  DialogFooter,
} from "@material-tailwind/react";
import "react-quill/dist/quill.snow.css";
import { getFileContent } from "@/libs/utils";
import {
  articleSave,
  articlesView,
  tagView,
  articleModify,
} from "@/libs/action";
import { XMarkIcon } from "@heroicons/react/24/solid";
import { useParams, useNavigate } from "react-router-dom";

export function Editor() {
  const navigate = useNavigate();
  const parapms = useParams();

  const [contentInfo, setContentInfo] = useState({
    title: "",
    content: "",
  });
  const [imageUploads, setImageUploads] = React.useState([]);

  const [outMaxFile, setOutMaxFile] = useState(false);
  const uploadImage = () => {
    let inputFile = document.createElement(`input`);
    inputFile.accept = "image/*";
    inputFile.type = "file";
    inputFile.click();
    inputFile.addEventListener("change", async () => {
      if (inputFile.files[0].size > 3 * 1024 * 1024) {
        setOutMaxFile(true);
        return;
      }

      if (inputFile.files && inputFile.files.length > 0) {
        let fileContent = await getFileContent(inputFile.files[0]);
        if (imageUploads.findIndex((item) => item == fileContent) == -1) {
          setImageUploads([...imageUploads, fileContent]);
        }
      }
    });
  };

  const [remoteTags, setRemoteTags] = useState([]);
  const [tagInfo, setTagInfo] = useState({
    tagInputType: true,
    tag: "",
    tags: [],
  });
  const remoteTagsClick = async () => {
    if (remoteTags.length != 0) return;
    const res = await tagView();
    if (res.status) {
      setRemoteTags(res.data);
    }
  };

  const [uploadStauts, setUploadStauts] = useState("");
  const uploadHandler = async () => {
    let id = parapms["id"];
    try {
      const res = id
        ? await articleModify(
            id,
            contentInfo.title,
            contentInfo.content,
            tagInfo.tags,
            imageUploads,
            true
          )
        : await articleSave(
            contentInfo.title,
            contentInfo.content,
            tagInfo.tags,
            imageUploads,
            true
          );

      if (res.status) setUploadStauts("success");
      else setUploadStauts("error");
    } catch (err) {
      setUploadStauts("error");
    }
  };

  const [resourceNotFound, setResourceNotFound] = useState(false);
  useEffect(() => {
    let id = parapms["id"];
    if (id) {
      articlesView(id)
        .then((res) => {
          if (!res.status) setResourceNotFound(true);
          const { Id, Title, Content, Covers, Tags } = res.data;
          setContentInfo({ title: Title, content: Content });
          setImageUploads(Covers);
          setTagInfo({ ...tagInfo, tags: Tags });
        })
        .catch(() => {
          setResourceNotFound(true);
        });
    }
  }, []);

  return (
    <div className="flex flex-col justify-center py-4 gap-y-4 gap-x-4">
      <div className="flex gap-4 items-center justify-between flex-wrap  xl:flex-col xl:items-start 4xl:flex-row">
        <div className="flex gap-4 w-full xl:w-1/2 flex-wrap 3xl:flex-nowrap">
          <div className="flex gap-x-4 gap-y-2 flex-wrap sm:flex-nowrap">
            <Input
              label="Post Title"
              value={contentInfo.title}
              onChange={(e) =>
                setContentInfo({ ...contentInfo, title: e.target.value })
              }
            />
            <ButtonGroup className="w-full ">
              <Button
                className="whitespace-nowrap"
                onClick={() => setTagInfo({ ...tagInfo, tagInputType: true })}
              >
                选择标签
              </Button>
              <Button
                className="whitespace-nowrap"
                onClick={() => setTagInfo({ ...tagInfo, tagInputType: false })}
              >
                输入标签
              </Button>
            </ButtonGroup>
          </div>
          <div className="flex gap-2 items-center flex-wrap 3xl:flex-nowrap">
            <div className="relative flex max-w-[24rem] flex-wrap">
              {tagInfo.tagInputType ? (
                <Select
                  onClick={remoteTagsClick}
                  label="Select Tag"
                  disabled={tagInfo.tags.length > 2}
                  className="w-56"
                >
                  {remoteTags.length == 0 ? (
                    <Option disabled>Empty Data</Option>
                  ) : (
                    remoteTags.map((item) => (
                      <Option
                        onClick={() => setTagInfo({ ...tagInfo, tag: item })}
                        key={item}
                      >
                        {item}
                      </Option>
                    ))
                  )}
                </Select>
              ) : (
                <Input
                  type="text"
                  label="input tag"
                  value={tagInfo.tag}
                  onChange={(v) => {
                    if (v.target.value.length > 12) return;
                    setTagInfo({ ...tagInfo, tag: v.target.value });
                  }}
                  className="w-56"
                  containerProps={{
                    className: "min-w-0",
                  }}
                />
              )}
              <Button
                size="sm"
                color={tagInfo.tag ? "gray" : "blue-gray"}
                disabled={!tagInfo.tag || tagInfo.tags.length > 2}
                className="!absolute right-1 top-1 rounded"
                onClick={() =>
                  setTagInfo({
                    ...tagInfo,
                    tags: [...tagInfo.tags, tagInfo.tag],
                  })
                }
              >
                添加
              </Button>
            </div>
            {tagInfo.tags.map((item, index) => (
              <Chip
                value={item}
                key={item + ":" + index}
                animate={{
                  mount: { y: 0 },
                  unmount: { y: 50 },
                }}
                onClose={() => {
                  setTagInfo({
                    ...tagInfo,
                    tags: tagInfo.tags.filter((i) => item != i),
                  });
                }}
              />
            ))}
          </div>
        </div>
        <div className="flex gap-2 flex-wrap flex-auto 4xl:justify-end">
          <div className="flex gap-x-4 flex-wrap">
            {imageUploads.map((item, index) => (
              <Badge
                content={
                  <XMarkIcon
                    onClick={() => {
                      setImageUploads(imageUploads.filter((i) => i != item));
                    }}
                    className="h-4 w-4 text-white  cursor-pointer"
                    strokeWidth={2.5}
                  />
                }
                className="bg-gradient-to-tr from-red-400 to-red-600 border-2 border-white shadow-lg shadow-black/20"
                key={index}
              >
                <img
                  src={item}
                  className="w-10 h-10 rounded-md object-cover"
                ></img>
              </Badge>
            ))}

            <Button
              className="whitespace-nowrap"
              disabled={!(imageUploads.length < 2)}
              onClick={uploadImage}
            >
              Upload Cover Image
            </Button>
          </div>
          <Button
            disabled={
              !contentInfo.title ||
              !contentInfo.content ||
              tagInfo.tags.length == 0
            }
          >
            暂存
          </Button>
          <Button
            onClick={uploadHandler}
            disabled={
              !contentInfo.title ||
              !contentInfo.content ||
              tagInfo.tags.length == 0
            }
          >
            上传
          </Button>
        </div>
      </div>
      <ReactQuill
        className="flex-auto h-[24rem]  lg:h-[62rem] border"
        theme="snow"
        value={contentInfo.content}
        onChange={(e) => setContentInfo({ ...contentInfo, content: e })}
      />
      <Dialog
        open={uploadStauts.length > 0}
        size="xs"
        handler={() => setUploadStauts("")}
      >
        <DialogHeader>
          上传{uploadStauts == "success" ? "成功" : "失败"}
        </DialogHeader>
        <DialogBody>
          上传
          {uploadStauts == "success" ? "成功, 请前去查看" : "失败, 请重新上传"}
        </DialogBody>
        <DialogFooter>
          <Button
            variant="gradient"
            color={uploadStauts == "success" ? "green" : "red"}
            onClick={() => setUploadStauts("")}
          >
            <span>知晓</span>
          </Button>
        </DialogFooter>
      </Dialog>
      <Dialog
        open={resourceNotFound}
        size="xs"
        handler={() => {
          setResourceNotFound(false);
          navigate("/dashboard/blogs");
        }}
      >
        <DialogHeader>资源不存在</DialogHeader>
        <DialogBody>你查找的资源不存在</DialogBody>
        <DialogFooter>
          <Button
            variant="gradient"
            color="red"
            onClick={() => {
              setResourceNotFound(false);
              navigate("/dashboard/blogs");
            }}
          >
            <span>知晓</span>
          </Button>
        </DialogFooter>
      </Dialog>
      <Dialog open={outMaxFile} size="xs" handler={() => setOutMaxFile(false)}>
        <DialogHeader>文件超出大小</DialogHeader>
        <DialogBody>选择的图片大小不应该大于3MB</DialogBody>
        <DialogFooter>
          <Button
            variant="gradient"
            color="red"
            onClick={() => setOutMaxFile(false)}
          >
            <span>知晓</span>
          </Button>
        </DialogFooter>
      </Dialog>
    </div>
  );
}

export default Editor;
