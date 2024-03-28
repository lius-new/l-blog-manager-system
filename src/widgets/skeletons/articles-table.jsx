import { Typography, Avatar, Chip } from "@material-tailwind/react";

const ArticlesTableSkeletions = () => {
  return (
    <>
      {Array.from("1234567890abcd").map((item) => {
        const className = `py-3 px-5 border-b border-blue-gray-50"}`;
        return (
          <tr key={item}>
            <td className={className}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
            <td className={className}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
            <td className={className}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
            <td className={className}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
            <td className={className}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
            <td className={className}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
            {/* TODO:  w-32 和 w-64: 该库当我设置w-32时会使button被隐藏 , 当我设置内部为64外部32就不会了*/}
            <td className={`${className} w-32`}>
              <Typography
                as="div"
                variant="paragraph"
                className="mb-2 h-6 w-full rounded-full bg-gray-300"
              >
                &nbsp;
              </Typography>
            </td>
          </tr>
        );
      })}
    </>
  );
};

export default ArticlesTableSkeletions;
