import {
  Card,
  CardHeader,
  CardBody,
  Typography,
  Avatar,
  Chip,
  CardFooter,
  IconButton,
  Button,
} from "@material-tailwind/react";
import {
  ArrowLeftIcon,
  ArrowRightIcon,
  InboxIcon,
} from "@heroicons/react/24/solid";
import { useEffect } from "react";
import { useState } from "react";
import {
  articleModify,
  articleModifyVisiable,
  articlesViews,
} from "@/libs/action";
import { useSearchParams } from "react-router-dom";
import ArticlesTableSkeletions from "@/widgets/skeletons/articles-table";
import { useNavigate } from "react-router-dom";

const dateFormat = (time) => {
  time = time + "000";
  let date = new Date(parseInt(time));

  return (
    date.getFullYear() +
    "/" +
    (date.getMonth() + 1) +
    "/" +
    date.getDate() +
    " " +
    date.getHours() +
    ":" +
    date.getMinutes() +
    ":" +
    date.getSeconds()
  );
};

export function Blogs() {
  const navigate = useNavigate();
  const pageNum = useSearchParams()[0].get("page");

  const [articles, setArticless] = useState([]);
  const [total, setTotal] = useState(1);
  const [pageList, setPageList] = useState([]);

  const geneartePageList = (n) => {
    let tempList = [];

    // n 是一共有多少条数据
    let pageCount = parseInt(n / 12) + 1;

    for (let index = 1; index <= pageCount; index++) {
      tempList.push(index);
    }
    return tempList;
  };

  const update = (n = 0) => {
    let pn = 1;
    if (n != 0) pn = n;
    else if (typeof pageNum == "string") navigate(`/dashboard/blogs?page=1`);
    else if (pageNum) pn = parseInt(pageNum[0]);

    // pageSize=12
    articlesViews(12, pn).then((res) => {
      if (res.code == 0 && res.message == "Ok") {
        setArticless(res.data.data);
        setTotal(res.data.total);
        setPageList([...geneartePageList(res.data.total)]);
      }
    });
  };

  useEffect(() => {
    update();
  }, []);

  const modifyHandle = (id, visiable) => {
    articleModifyVisiable(id, visiable).then((res) => {
      if (res.code == 0 && res.message == "Ok") {
        update();
      }
    });
  };
  const editorHandle = (id) => {
    navigate(`/dashboard/editor/${id}`);
  };

  return (
    <div className="mt-10 mb-8 flex flex-col gap-12">
      <Card>
        <CardHeader variant="gradient" color="gray" className="mb-8 p-6">
          <Typography variant="h6" color="white">
            文章管理
          </Typography>
        </CardHeader>
        <CardBody className="overflow-x-scroll px-0 pt-0 pb-2 overflow-hidden">
          {articles == null ? (
            <div className="flex flex-col items-center justify-center w-full">
              <InboxIcon className="w-12 h-12" />
              data empty
            </div>
          ) : (
            <table className="w-full min-w-[640px] table-auto ">
              <thead>
                <tr className="grid grid-cols-7 ">
                  {["编号", "图片", "标题", "标签", "状态", "时间", "操作"].map(
                    (el) => (
                      <th
                        key={el}
                        className="border-b border-blue-gray-50 py-3 px-5 text-left"
                      >
                        <Typography
                          variant="small"
                          className="text-[11px] font-bold uppercase text-blue-gray-400"
                        >
                          {el}
                        </Typography>
                      </th>
                    )
                  )}
                </tr>
              </thead>
              <tbody>
                {articles.length > 0 ? (
                  articles.map(
                    ({ id, title, tags, visiable, updateAt, covers }) => {
                      const className = `py-3 px-5 border-b border-blue-gray-50"`;
                      return (
                        <tr key={id} className="grid grid-cols-7 ">
                          <td className={`${className} col-span-1`}>
                            <div className="flex items-center gap-4">{id}</div>
                          </td>
                          <td className={`${className} col-span-1`}>
                            <div className="flex items-center gap-4">
                              {covers &&
                                covers.map((item) => (
                                  <Avatar
                                    key={item}
                                    src={`${
                                      import.meta.env
                                        .VITE_CONTENT_API_SERVER_URI
                                    }/articles/image/${item}`}
                                    alt={item}
                                    size="sm"
                                    variant="rounded"
                                  />
                                ))}
                            </div>
                          </td>
                          <td className={`${className} col-span-1`}>
                            <div>
                              <Typography
                                variant="small"
                                color="blue-gray"
                                className="font-semibold"
                              >
                                {title}
                              </Typography>
                            </div>
                          </td>
                          <td className={className}>
                            {tags &&
                              tags.map((item) => (
                                <span
                                  className="text-xs font-semibold text-blue-gray-800 bg-gray-300 px-2 py-1 rounded-md"
                                  key={item}
                                >
                                  {item}
                                </span>
                              ))}
                          </td>
                          <td className={className}>
                            <Chip
                              variant="gradient"
                              color={visiable ? "green" : "blue-gray"}
                              value={visiable ? "online" : "offline"}
                              className="py-0.5 px-2 text-[11px] h-8 font-medium w-20 text-center"
                            />
                          </td>
                          <td className={`${className} `}>
                            <Typography className="text-xs font-semibold text-blue-gray-600">
                              {dateFormat(updateAt)}
                            </Typography>
                          </td>
                          {/* TODO:  w-32 和 w-64: 该库当我设置w-32时会使button被隐藏 , 当我设置内部为64外部32就不会了*/}
                          <td className={`${className}`}>
                            <div className="flex gap-x-2 ">
                              <Button
                                size="sm"
                                color="amber"
                                onClick={() => editorHandle(id)}
                              >
                                Edit
                              </Button>
                              <Button
                                size="sm"
                                color={visiable ? "red" : "green"}
                                className="w-20"
                                onClick={() => modifyHandle(id, !visiable)}
                              >
                                {visiable ? "Disable" : "Enable"}
                              </Button>
                            </div>
                          </td>
                        </tr>
                      );
                    }
                  )
                ) : (
                  <ArticlesTableSkeletions />
                )}
              </tbody>
            </table>
          )}
        </CardBody>

        <CardFooter>
          <div className="flex items-center gap-4 justify-end">
            <Button
              variant="text"
              className="flex items-center gap-2 rounded-full"
              disabled={pageNum == 1 || !pageNum}
              onClick={() => {
                navigate(`/dashboard/blogs?page=${pageNum - 1}`);
                update(pageNum - 1);
              }}
            >
              <ArrowLeftIcon strokeWidth={2} className="h-4 w-4" /> Previous
            </Button>
            <div className="flex items-center gap-2">
              {pageList.map((item) => (
                <IconButton
                  className={`${
                    pageNum == item || (!pageNum && item == 1)
                      ? "bg-blue-gray-600"
                      : "bg-blue-gray-300"
                  }`}
                  key={item}
                  onClick={() => {
                    navigate(`/dashboard/blogs?page=${item}`);
                    update(item);
                  }}
                >
                  {item}
                </IconButton>
              ))}
            </div>
            <Button
              variant="text"
              className="flex items-center gap-2 rounded-full"
              disabled={pageNum == pageList.length}
              onClick={() => {
                if (!pageNum) {
                  navigate(`/dashboard/blogs?page=2`);
                  update(2);
                } else {
                  navigate(`/dashboard/blogs?page=${parseInt(pageNum) + 1}`);
                  update(parseInt(pageNum) + 1);
                }
              }}
            >
              Next
              <ArrowRightIcon strokeWidth={2} className="h-4 w-4" />
            </Button>
          </div>
        </CardFooter>
      </Card>
    </div>
  );
}

export default Blogs;
