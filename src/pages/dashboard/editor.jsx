import React, { useEffect, useState } from "react";
import {
  Switch,
  Input,
  Select,
  Option,
  Button,
  Chip,
  Badge,
  Textarea,
  Dialog,
  DialogHeader,
  DialogBody,
  DialogFooter,
} from "@material-tailwind/react";
import "react-quill/dist/quill.snow.css";
import {
  base64ToImageUrl,
  getContentImage,
  getFileContent,
  getFileImageContent,
} from "@/libs/utils";
import {
  articleSave,
  articlesView,
  tagView,
  articleModify,
  uploadArticleInnerImages,
  coverViewBase64,
} from "@/libs/action";
import { BanknotesIcon, PlusIcon, XMarkIcon } from "@heroicons/react/24/solid";
import { useParams, useNavigate } from "react-router-dom";
import { MdEditor, NormalToolbar } from "md-editor-rt";
import "md-editor-rt/lib/style.css";

export function Editor() {
  const navigate = useNavigate();
  const parapms = useParams();

  const [visiable, setVisiable] = useState(false);
  const [contentInfo, setContentInfo] = useState({
    title: "",
    content: "",
    description: "",
  });

  // 上传的图片文件,临时显示
  const [imageShowUploads, setImageShowUploads] = React.useState([]);

  const [outMaxFile, setOutMaxFile] = useState(false);
  const uploadImage = () => {
    let inputFile = document.createElement(`input`);
    inputFile.accept = "image/*";
    inputFile.type = "file";
    inputFile.click();
    inputFile.addEventListener("change", async () => {
      if (inputFile.files[0].size > 3 * 1024 * 1024) {
        // 判断文件大小
        setOutMaxFile(true);
        return;
      }

      if (inputFile.files && inputFile.files.length > 0) {
        let fileContent = await getFileImageContent(inputFile.files[0]);
        console.log(fileContent);
        if (imageShowUploads.findIndex((item) => item == fileContent) == -1) {
          setImageShowUploads([...imageShowUploads, fileContent]);
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
    if (res.code == 0 && res.message == "Ok") {
      setRemoteTags(res.data.data.map((item) => item.name));
    }
  };
  const uploadFileHandle = () => {
    let inputFile = document.createElement(`input`);
    inputFile.accept = ".md";
    inputFile.type = "file";
    inputFile.click();
    inputFile.addEventListener("change", async () => {
      if (inputFile.files && inputFile.files.length > 0) {
        let fileContent = await getFileContent(inputFile.files[0]);
        setContent(content + fileContent);
      }
    });
  };

  const [contentImages, setContentImages] = useState([]);
  const markdownUpliadImage = async (files, callback) => {
    const res = await Promise.all(
      files.map(async (file) => {
        return new Promise(async (rev, rej) => {
          const imageBase64 = await getFileImageContent(file);
          const url = base64ToImageUrl(imageBase64);
          rev(url);
          setContentImages([
            ...contentImages,
            { key: url, value: imageBase64 },
          ]);
        });
      })
    );

    callback(res.map((item) => item));
  };

  const [uploadStauts, setUploadStauts] = useState("");
  const uploadHandler = async () => {
    let id = parapms["id"];

    // 修改content中的images
    // 因为 contentImages 实际上在每次插入图片都会添加(MdEditor.onUploadImg插入图片时触发，删除图片时不触发), 所以当删除时仍然有图片信息在contentImages中。
    // content state中包含的![]()才是实际上的图片. cImages通过正则表达式获取到的所有的图片(blob url), 其中的每个对应content state中某个图片的key(blob url). 如果对应那么就说明图片是存在在content中的.
    let cImages = getContentImage(content); // markdown中所有的![]()中的链接部分(`()`中的部分)

    // 过滤掉contentImages中被删除的部分(因为删除不触发所以手动过滤)
    // 如果和外部的image图片，那么肯定不会去调用MdEditor.onUploadImg方法，同样也会被过滤掉.
    let contentImagesFilter = contentImages.filter(
      (item) => cImages.findIndex((i) => i.includes(item.key)) != -1
    );

    // 先上传图片,更新contentInfo.content中所有原本blog url换成服务器实际上图片的地址
    // TODO: 目前的image是放到响应体中会导致请求实体过大而无法上传很大的图片文件。
    let tempContent = JSON.parse(JSON.stringify(content));
    if (contentImagesFilter.length > 0) {
      let contentImageUploadRes = await uploadArticleInnerImages(
        contentImagesFilter.map((item) => item.value)
      );

      if (
        contentImageUploadRes.code == 0 &&
        contentImageUploadRes.message == "Ok"
      ) {
        let newContentImages = contentImageUploadRes.data.hashs.map(
          (item, index) => {
            contentImagesFilter[index]["newKey"] = `${
              import.meta.env.VITE_CONTENT_API_SERVER_URI
            }/articles/image/${item}`;
            return contentImagesFilter[index];
          }
        );
        newContentImages.forEach((element) => {
          tempContent = tempContent.replace(element.key, element.newKey);
        });
      }
    }
    if (tempContent != content) {
      setContent(tempContent);
    }

    try {
      let res = id
        ? await articleModify(
            id,
            contentInfo.title,
            tempContent,
            contentInfo.description,
            tagInfo.tags,
            imageShowUploads,
            visiable
          )
        : await articleSave(
            contentInfo.title,
            tempContent,
            contentInfo.description,
            tagInfo.tags,
            imageShowUploads
          );

      if (res.code == 0 && res.message == "Ok") {
        setUploadStauts("操作成功");
      } else {
        setUploadStauts(res.message);
      }
    } catch (err) {
      console.log(err);
      setUploadStauts("error");
    }
  };

  const [resourceNotFound, setResourceNotFound] = useState(false);
  useEffect(() => {
    let id = parapms["id"];
    if (id) {
      articlesView(id)
        .then((res) => {
          if (res.code != 0 || res.message != "Ok") setResourceNotFound(true);
          const { title, content, description, covers, tags, visiable } =
            res.data.data;
          setContentInfo({
            title: title,
            description: description,
          });
          setContent(content);
          if (covers) {
            covers.forEach((item, index) => {
              coverViewBase64(item).then((res) => {
                let tempImageShowUploads = JSON.parse(
                  JSON.stringify(imageShowUploads)
                );
                tempImageShowUploads[index] = res;
                setImageShowUploads(tempImageShowUploads);
              });
            });
          }
          setTagInfo({ ...tagInfo, tags: tags });
          setVisiable(visiable);
        })
        .catch(() => {
          setResourceNotFound(true);
        });
    }
  }, []);

  const [settingOpen, setSettingOpen] = useState(false);
  const [content, setContent] = useState("");

  return (
    <div className="flex flex-col justify-center py-4 gap-y-4 gap-x-4">
      <div className="flex gap-x-4 justify-end">
        <Button
          onClick={() => {
            setSettingOpen(true);
          }}
        >
          文章属性
        </Button>
        <Button
          onClick={uploadHandler}
          disabled={!contentInfo.title || !content || tagInfo.tags.length == 0}
        >
          上传
        </Button>
      </div>

      <MdEditor
        codeTheme={"atom"}
        modelValue={content}
        onChange={(e) => setContent(e)}
        onUploadImg={markdownUpliadImage}
        className="flex-auto h-[24rem]  lg:h-[60rem] border rounded-md"
        toolbars={[
          "bold",
          "italic",
          "underline",
          "-",
          "strikeThrough",
          "title",
          "sub",
          "sup",
          "quote",
          "unorderedList",
          "orderedList",
          "codeRow",
          "code",
          "link",
          "image",
          "table",
          "mermaid",
          "katex",
          "task",
          0,
          "=",
          "revoke",
          "next",
          "prettier",
          "pageFullscreen",
          "fullscreen",
          "preview",
          "htmlPreview",
          "catalog",
          "github",
        ]}
      />
      <SettingModal
        title={"设置文章信息"}
        open={settingOpen}
        handle={() => setSettingOpen(false)}
        contentInfo={contentInfo}
        setContentInfo={setContentInfo}
        tagInfo={tagInfo}
        setTagInfo={setTagInfo}
        uploadImage={uploadImage}
        imageShowUploads={imageShowUploads}
        setImageShowUploads={setImageShowUploads}
        remoteTags={remoteTags}
        remoteTagsClick={remoteTagsClick}
      />
      <MessageModal
        title={"Message"}
        content={uploadStauts}
        open={uploadStauts.length > 0}
        handle={() => {
          setUploadStauts("");
        }}
      />
      <MessageModal
        title={"资源不存在"}
        content={"你查找的资源不存在"}
        open={resourceNotFound}
        handle={() => {
          setResourceNotFound(false);
          navigate("/dashboard/blogs");
        }}
      />
      <MessageModal
        title={"文件超出大小"}
        content={"选择的图片大小不应该大于3MB"}
        open={outMaxFile}
        handle={() => setOutMaxFile(false)}
      />
    </div>
  );
}

const MessageModal = ({ title, content, open, handle }) => {
  return (
    <Dialog open={open} size="xs" handler={handle}>
      <DialogHeader>{title}</DialogHeader>
      <DialogBody>{content}</DialogBody>
      <DialogFooter>
        <Button variant="gradient" onClick={handle}>
          <span>知晓</span>
        </Button>
      </DialogFooter>
    </Dialog>
  );
};

const SettingModal = ({
  title,
  open,
  handle,
  contentInfo,
  setContentInfo,
  tagInfo,
  setTagInfo,
  uploadImage,
  imageShowUploads,
  setImageShowUploads,
  remoteTags,
  remoteTagsClick,
}) => {
  return (
    <Dialog open={open} size={"sm"} handler={handle}>
      <DialogHeader className="flex justify-center">
        <h3 className="text-center text-gray-700">{title}</h3>
      </DialogHeader>
      <DialogBody>
        <div className="flex gap-x-4 flex-col ">
          <div className="flex flex-col gap-y-4 text-sm">
            <div className="flex flex-col gap-y-2">
              <span>文章标题</span>
              <Input
                label="Post Title"
                value={contentInfo.title}
                onChange={(e) =>
                  setContentInfo({ ...contentInfo, title: e.target.value })
                }
              />
            </div>
            <div className="flex flex-col gap-y-2">
              <div className="flex justify-between items-center">
                <span>文章标签</span>
                <Switch
                  onChange={() =>
                    setTagInfo({
                      ...tagInfo,
                      tagInputType: !tagInfo.tagInputType,
                    })
                  }
                />
              </div>
              <div className="flex items-center gap-x-2">
                {tagInfo.tagInputType ? (
                  <Select
                    onFocus={remoteTagsClick}
                    label="Select Tag"
                    disabled={tagInfo.tags.length > 2}
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
                  <>
                    <Input
                      type="text"
                      label="input tag"
                      value={tagInfo.tag}
                      onChange={(v) => {
                        if (v.target.value.length > 12) return;
                        setTagInfo({ ...tagInfo, tag: v.target.value });
                      }}
                    />
                  </>
                )}
                <button
                  disabled={!tagInfo.tag || tagInfo.tags.length > 2}
                  className={`w-14 h-8 rounded p-2 py-1.5 ${
                    tagInfo.tag
                      ? "cursor-pointer bg-gray-900 text-white"
                      : "cursor-not-allowed bg-gray-300"
                  }`}
                  onClick={() =>
                    setTagInfo({
                      ...tagInfo,
                      tags: [...tagInfo.tags, tagInfo.tag],
                      tag: "",
                    })
                  }
                >
                  添加
                </button>
              </div>

              <div className="flex gap-x-2 flex-wrap">
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
          </div>
          <div className="mt-4">
            <span>文章描述</span>
            <Textarea
              label=""
              className="!border !border-gray-300 bg-white text-gray-900 shadow-lg shadow-gray-900/5 ring-4 ring-transparent placeholder:text-gray-500 placeholder:opacity-100 focus:!border-gray-900 focus:!border-t-gray-900 focus:ring-gray-900/10"
              labelProps={{
                className: "hidden",
              }}
              placeholder="请输入简短的文章描述信息..."
              value={contentInfo.description}
              onChange={(e) =>
                setContentInfo({ ...contentInfo, description: e.target.value })
              }
            />
          </div>
        </div>
        <hr className="my-4" />
        <div className="mb-4 flex flex-col gap-2 text-sm select-none">
          <span>添加图片</span>
          {Array.from("12").map((item, index) => {
            let imageShowUpload = imageShowUploads[index];
            if (imageShowUpload) {
              return (
                <Badge
                  content={
                    <XMarkIcon
                      onClick={() => {
                        setImageShowUploads(
                          imageShowUploads.filter((i) => i != imageShowUpload)
                        );
                      }}
                      className="h-4 w-4 text-white  cursor-pointer"
                      strokeWidth={2.5}
                    />
                  }
                  className="bg-gradient-to-tr from-red-400 to-red-600 border-2 border-white shadow-lg shadow-black/20"
                  key={index}
                >
                  <img
                    src={imageShowUpload}
                    className="w-full h-40 rounded-md object-cover object-center"
                  ></img>
                </Badge>
              );
            } else {
              return (
                <Button
                  key={item}
                  variant="outlined"
                  className="flex justify-center items-center w-full h-40 border-gray-400 rounded-md"
                  onClick={uploadImage}
                >
                  <PlusIcon className="text-xs w-10"></PlusIcon>
                </Button>
              );
            }
          })}
        </div>
      </DialogBody>
    </Dialog>
  );
};
export default Editor;
