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
import { articlesView } from "@/libs/action";
import { useSearchParams, NavLink } from "react-router-dom";
import ArticlesTableSkeletions from "@/widgets/skeletons/articles-table";
import { useNavigate } from "react-router-dom";

export function Blogs() {
  const navigate = useNavigate();
  const pageNum = useSearchParams()[0].get("page");

  const [articles, setArticless] = useState([]);
  const [total, setTotal] = useState(1);
  const [pageList, setPageList] = useState([]);

  const geneartePageList = (n) => {
    let tempList = [];

    for (let index = 1; index <= n; index++) {
      tempList.push(index);
    }
    return tempList;
  };

  const update = (n = 0) => {
    let pn = 1;
    if (n != 0) pn = n;
    else if (pageNum) pn = parseInt(pageNum[0]);

    articlesView(14, pn).then((res) => {
      if (res.status) {
        setArticless(res.data);
        setTotal(res.total);
        setPageList([...geneartePageList(res.total)]);
      }
    });
  };

  useEffect(() => {
    update();
  }, []);

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
                <tr>
                  {["", "标题", "内容", "标签", "状态", "时间", "操作"].map(
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
                    ({ Id, Title, Content, Tags, Status, Time, Covers }) => {
                      const className = `py-3 px-5 border-b border-blue-gray-50"`;
                      return (
                        <tr key={Id}>
                          <td className={className}>
                            <div className="flex items-center gap-4">
                              {Covers.map((item) => (
                                <Avatar
                                  key={item}
                                  src={item}
                                  alt={item}
                                  size="sm"
                                  variant="rounded"
                                />
                              ))}
                            </div>
                          </td>
                          <td className={className}>
                            <div>
                              <Typography
                                variant="small"
                                color="blue-gray"
                                className="font-semibold"
                              >
                                {Title}
                              </Typography>
                            </div>
                          </td>
                          <td className={className}>
                            <div>
                              <Typography
                                variant="small"
                                color="blue-gray"
                                className="font-semibold"
                              >
                                {Content}
                              </Typography>
                            </div>
                          </td>
                          <td className={className}>
                            <Typography className="text-xs font-semibold text-blue-gray-600">
                              {Tags}
                            </Typography>
                            <Typography className="text-xs font-normal text-blue-gray-500">
                              {Tags}
                            </Typography>
                          </td>
                          <td className={className}>
                            <Chip
                              variant="gradient"
                              color={Status ? "green" : "blue-gray"}
                              value={Status ? "online" : "offline"}
                              className="py-0.5 px-2 text-[11px] font-medium w-fit"
                            />
                          </td>
                          <td className={className}>
                            <Typography className="text-xs font-semibold text-blue-gray-600">
                              {Time}
                            </Typography>
                          </td>
                          {/* TODO:  w-32 和 w-64: 该库当我设置w-32时会使button被隐藏 , 当我设置内部为64外部32就不会了*/}
                          <td className={`${className} w-32`}>
                            <div className="w-64 flex gap-x-2">
                              <Button size="sm" color="blue">
                                Change
                              </Button>
                              <Button size="sm" color="amber">
                                Edit
                              </Button>
                              <Button size="sm" color="red">
                                Disable
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
            >
              <ArrowLeftIcon strokeWidth={2} className="h-4 w-4" /> Previous
            </Button>
            <div className="flex items-center gap-2">
              {pageList.map((item) => (
                <IconButton
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
