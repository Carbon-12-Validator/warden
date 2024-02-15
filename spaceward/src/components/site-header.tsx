import { Icons } from "@/components/icons";
import { ThemeToggle } from "@/components/theme-toggle";
import { Notifications } from "@/components/notifications";
import { WalletConnect } from "@/components/walletconnect";
import { ConnectWallet } from "@/components/connect-wallet";

export function SiteHeader() {
	return (
		<header className="bg-background fixed top-0 z-[60] w-full border-b">
			<div className="pl-6 flex h-16 items-center space-x-4 sm:justify-between sm:space-x-0">
				<a
					href="/"
					className="flex items-center space-x-2"
				>
					<Icons.logo className="h-6 w-auto" />
				</a>
				<div className="flex flex-1 items-center justify-end space-x-4">
					<nav className="flex items-center">
						{/* <ChainsDropdown /> */}
						<Notifications />
						<WalletConnect />
						<ConnectWallet />
						<ThemeToggle />
					</nav>
				</div>
			</div>
		</header>
	);
}