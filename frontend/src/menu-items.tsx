import { MenuItem } from "@/types/types";

export const leftMenuItems: MenuItem[] = [
  {
    id: "dashboard",
    icon: "NiHome",
    label: "Dashboard",
    description: "Main dashboard",
    color: "text-primary",
    href: "/dashboard",
  },
  {
    id: "curriculum",
    icon: "NiBook",
    label: "Curriculum",
    description: "Curriculum management",
    color: "text-primary",
    href: "/curriculum",
    children: [
      {
        id: "cp",
        icon: "NiDocumentFull",
        label: "CP",
        href: "/cp",
        description: "Curriculum Plan",
      },
      {
        id: "tp",
        icon: "NiDocumentFull",
        label: "TP",
        href: "/tp",
        description: "Teaching Plan",
      },
      {
        id: "atp",
        icon: "NiDocumentFull",
        label: "ATP",
        href: "/atp",
        description: "Annual Teaching Plan",
      },
    ],
  },
  {
    id: "learning-design",
    icon: "NiGraduation",
    label: "Learning Design",
    description: "Learning design and materials",
    color: "text-primary",
    href: "/learning-design",
    children: [
      {
        id: "modul-ajar",
        icon: "NiBook",
        label: "Modul Ajar",
        href: "/modul-ajar",
        description: "Teaching modules",
      },
    ],
  },
  {
    id: "assessment",
    icon: "NiListCheck",
    label: "Assessment",
    description: "Assessment management",
    color: "text-primary",
    href: "/assessment",
    children: [
      {
        id: "assessment-list",
        icon: "NiListCheck",
        label: "Assessment",
        href: "/assessment",
        description: "Assessment list",
      },
      {
        id: "rubric",
        icon: "NiStars",
        label: "Rubric",
        href: "/rubric",
        description: "Assessment rubrics",
      },
      {
        id: "narrative-report",
        icon: "NiDocumentFull",
        label: "Narrative Report",
        href: "/narrative-report",
        description: "Narrative reports",
      },
    ],
  },
  {
    id: "workflow",
    icon: "NiPath",
    label: "Workflow",
    description: "Workflow and approvals",
    color: "text-primary",
    href: "/workflow",
    children: [
      {
        id: "approval-queue",
        icon: "NiCheckSquare",
        label: "Approval Queue",
        href: "/workflow",
        description: "Pending approvals",
      },
    ],
  },
];

export const leftMenuBottomItems: MenuItem[] = [
  {
    id: "administration",
    icon: "NiSettings",
    label: "Administration",
    href: "/settings",
    color: "text-primary",
    children: [
      {
        id: "settings",
        icon: "NiSettings",
        label: "Settings",
        href: "/settings",
        description: "System settings",
      },
    ],
  },
];
