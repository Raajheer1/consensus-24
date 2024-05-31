import type { Link } from "@/types";

const SidebarLinks: Link[] = [
  {
    title: "Home",
    icon: "fa-solid fa-home",
    to: { name: "Home" },
  },
  {
    title: "Ramp",
    icon: "fa-solid fa-money-bill-transfer",
    to: { name: "Ramp" },
  },
  {
    title: "Loan Overview",
    icon: "fa-solid fa-landmark",
    subLinks: [
      {
        title: "Available Loans",
        icon: "fa-solid fa-folder",
        to: { name: "Loans" },
      },
    ],
    showSubLinks: true,
  },
  {
    title: "My Profile",
    icon: "fa-solid fa-user",
    to: { name: "Profile" },
  },
];

export default SidebarLinks;
