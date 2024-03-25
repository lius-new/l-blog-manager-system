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
  ButtonGroup,
} from "@material-tailwind/react";
import { authorsTableData } from "@/data";
import { ArrowLeftIcon, ArrowRightIcon } from "@heroicons/react/24/solid";

export function Blogs() {
  return (
    <div className="mt-10 mb-8 flex flex-col gap-12">
      <Card>
        <CardHeader variant="gradient" color="gray" className="mb-8 p-6">
          <Typography variant="h6" color="white">
            文章管理
          </Typography>
        </CardHeader>
        <CardBody className="overflow-x-scroll px-0 pt-0 pb-2 overflow-hidden">
          <table className="w-full min-w-[640px] table-auto ">
            <thead>
              <tr>
                {["标题", "标签", "状态", "统计", "操作"].map((el) => (
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
                ))}
              </tr>
            </thead>
            <tbody>
              {authorsTableData.map(
                ({ img, name, email, job, online, date }, key) => {
                  const className = `py-3 px-5 ${
                    key === authorsTableData.length - 1
                      ? ""
                      : "border-b border-blue-gray-50"
                  }`;

                  return (
                    <tr key={name}>
                      <td className={className}>
                        <div className="flex items-center gap-4">
                          <Avatar
                            src={img}
                            alt={name}
                            size="sm"
                            variant="rounded"
                          />
                          <div>
                            <Typography
                              variant="small"
                              color="blue-gray"
                              className="font-semibold"
                            >
                              {name}
                            </Typography>
                            <Typography className="text-xs font-normal text-blue-gray-500">
                              {email}
                            </Typography>
                          </div>
                        </div>
                      </td>
                      <td className={className}>
                        <Typography className="text-xs font-semibold text-blue-gray-600">
                          {job[0]}
                        </Typography>
                        <Typography className="text-xs font-normal text-blue-gray-500">
                          {job[1]}
                        </Typography>
                      </td>
                      <td className={className}>
                        <Chip
                          variant="gradient"
                          color={online ? "green" : "blue-gray"}
                          value={online ? "online" : "offline"}
                          className="py-0.5 px-2 text-[11px] font-medium w-fit"
                        />
                      </td>
                      <td className={className}>
                        <Typography className="text-xs font-semibold text-blue-gray-600">
                          {date}
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
              )}
            </tbody>
          </table>
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
              <IconButton>1</IconButton>
              <IconButton>2</IconButton>
              <IconButton>3</IconButton>
              <IconButton>4</IconButton>
              <IconButton>5</IconButton>
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
