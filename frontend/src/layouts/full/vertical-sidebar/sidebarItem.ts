import {
    ApertureIcon,
    CopyIcon,
    LayoutDashboardIcon,
    LoginIcon,
    MoodHappyIcon,
    TypographyIcon,
    UserPlusIcon,
    FileUploadIcon,
    LivePhotoIcon
} from 'vue-tabler-icons';

export interface menu {
    header?: string;
    title?: string;
    icon?: any;
    to?: string;
    chip?: string;
    chipColor?: string;
    chipVariant?: string;
    chipIcon?: string;
    children?: menu[];
    disabled?: boolean;
    type?: string;
    subCaption?: string;
}

const sidebarItem: menu[] = [
    {
        title: 'Dashboard',
        icon: ApertureIcon,
        to: '/'
    },
    { header: 'Home' },
    {
        title: '网易云音乐转换',
        icon: FileUploadIcon,
        to: '/netease/index'
    },
    {
        title: 'BING图片',
        icon: LivePhotoIcon,
        to: '/bing/images'
    }
];

export default sidebarItem;
