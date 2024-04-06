import React, { useEffect, useState } from "react";
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
import { base64ToImageUrl, getContentImage, getFileContent, getFileImageContent } from "@/libs/utils";
import {
  articleSave,
  articlesView,
  tagView,
  articleModify,
  uploadArticleInnerImages,
} from "@/libs/action";
import { BanknotesIcon, XMarkIcon } from "@heroicons/react/24/solid";
import { useParams, useNavigate } from "react-router-dom";
import { MdEditor, NormalToolbar } from 'md-editor-rt';
import 'md-editor-rt/lib/style.css';

export function Editor() {
  const navigate = useNavigate();
  const parapms = useParams();

  const [contentInfo, setContentInfo] = useState({
    title: "",
    content: "",
    description: "",
  });

  // 上传的图片文件,提交到后台
  const [imageUploads, setImageUploads] = React.useState([]);
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
        setOutMaxFile(true);
        return;
      }

      if (inputFile.files && inputFile.files.length > 0) {
        let fileContent = await getFileImageContent(inputFile.files[0]);
        if (imageUploads.findIndex((item) => item == fileContent) == -1) {
          setImageShowUploads([...imageShowUploads, fileContent]);
        }
        setImageUploads([...imageUploads, inputFile.files[0]]);
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
  const uploadFileHandle = () => {
    let inputFile = document.createElement(`input`);
    inputFile.accept = ".md";
    inputFile.type = "file";
    inputFile.click();
    inputFile.addEventListener("change", async () => {
      if (inputFile.files && inputFile.files.length > 0) {
        let fileContent = await getFileContent(inputFile.files[0]);
        setContentInfo({ ...contentInfo, content: contentInfo.content + fileContent })
      }
    });
  }

  const [contentImages, setContentImages] = useState([])
  const markdownUpliadImage = async (files, callback) => {
    const res = await Promise.all(
      files.map(async (file) => {
        return new Promise(async (rev, rej) => {
          const imageBase64 = await getFileImageContent(file)
          const url = base64ToImageUrl(imageBase64)
          rev(url)
          setContentImages([...contentImages, { key: url, value: file }])
        });
      })
    );

    callback(res.map((item) => item));
  }

  const [uploadStauts, setUploadStauts] = useState("");
  const uploadHandler = async () => {
    let id = parapms["id"];

    // 修改content中的images, 因为 contentImages 实际上在每次插入图片都会添加,但是删除就不删除. 而contentInfo中包含的![]()才是实际上的图片. cImages通过正则表达式获取到的所有的图片(blob url), 其中的每个对应contentImages中某个图片的key(blob url). 如果对应那么就说明图片是存在在content中的.
    let cImages = getContentImage(contentInfo.content)
    let contentImage = contentImages.filter(item =>
      cImages.findIndex(i => i.includes(item.key)) != -1
    )



    // 先上传图片,更新contentInfo.content中所有原本blog url换成服务器实际上图片的地址
    if (contentImage.length > 0) {
      let contentImageUploadRes = await uploadArticleInnerImages(contentImage.map(item => item.value))
      contentImageUploadRes = contentImageUploadRes.data.map((item, index) => { contentImage[index]["newKey"] = `${import.meta.env.VITE_API_SERVER_URI}/api/file/${item}`; return contentImage[index] })
      contentImageUploadRes.forEach(element => {
        setContentInfo({ title: contentInfo.title, content: contentInfo.content.replace(element.key, element.newKey) })
      });
    }

    let description = contentInfo.content.match(/--!\s*\w+/g)
    description = description.length > 0 ? description[0] : ""

    try {
      const res = id
        ? await articleModify(
          id,
          contentInfo.title,
          contentInfo.content.replace(description, ""),
          description.replace("--!", ""),
          tagInfo.tags,
          imageUploads,
        )
        : await articleSave(
          contentInfo.title,
          contentInfo.content.replace(description, ""),
          description.replace("--!", ""),
          tagInfo.tags,
          imageUploads,
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
          const { Id, Title, Content, Description, Covers, Tags } = res.data;
          console.log(Covers);
          setContentInfo({ title: Title, content: "--!" + Description + "\n" + Content });
          setImageUploads(Covers);
          setImageShowUploads(Covers);
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
            {imageShowUploads.map((item, index) => (
              <Badge
                content={
                  <XMarkIcon
                    onClick={() => {
                      setImageShowUploads(imageShowUploads.filter((i) => i != item));
                      imageUploads.splice(index, 1)
                      setImageUploads(imageUploads);
                    }}
                    className="h-4 w-4 text-white  cursor-pointer"
                    strokeWidth={2.5}
                  />
                }
                className="bg-gradient-to-tr from-red-400 to-red-600 border-2 border-white shadow-lg shadow-black/20"
                key={index}
              >
                <img
                  src={item.includes('data:image/') ? item : `${import.meta.env.VITE_API_SERVER_URI}/api/file/${item}`}
                  className="w-10 h-10 rounded-md object-cover"
                ></img>
              </Badge>
            ))}

            <Button
              className="whitespace-nowrap"
              onClick={uploadFileHandle}
            >
              读取文章
            </Button>
            <Button
              className="whitespace-nowrap"
              disabled={!(imageShowUploads.length < 2)}
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

      <MdEditor
        codeTheme={"atom"}
        modelValue={contentInfo.content}
        onChange={(e) => setContentInfo({ ...contentInfo, content: e })}
        onUploadImg={markdownUpliadImage}
        className="flex-auto h-[24rem]  lg:h-[60rem] border rounded-md"
        toolbars={['bold', 'italic', 'underline', '-', "strikeThrough", "title", "sub", "sup", "quote", 'unorderedList',
          'orderedList', 'codeRow', 'code', 'link', 'image', 'table', 'mermaid', 'katex', 'task', 0,
          '=', 'revoke', 'next', 'prettier', 'pageFullscreen', 'fullscreen', 'preview', 'htmlPreview', 'catalog', 'github']}
        defToolbars={[
          <NormalToolbar
            title="标题"
            onClick={() => setContentInfo({ ...contentInfo, content: contentInfo.content + "\n--! " })}
            trigger={
              <BanknotesIcon className="w-4 h-4" />
            }
          />
        ]}
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
